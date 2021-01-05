// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/errors"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/compare"
)

const (
	contextCompareToValueKey    = "CompareToValue"
	contextIsComparableTypesKey = "IsComparable"
)

var _ interfaces.Validator = &compareValidator{}

type compareFunc func(ctx *ValidatorContext) (result bool, err error)

// compareValidator provides common functionality for all comparison validators (equal, greaterThan, lessThan, etc)
type compareValidator struct {
	*AllFieldValidators
	validatorType interfaces.Validator
	test          compareFunc
}

type valueCompare struct {
	getValue            interfaces.ValueFromContext
	compareToContextKey string
	comparisonValues    []int
}

func (v *valueCompare) evaluate(ctx *ValidatorContext) (bool, error) {
	contextData := ctx.GetContextData()
	comparer := compare.NewDefaultComparer(contextData[ContextFieldValueKey])
	c, err := comparer.Compare(contextData[v.compareToContextKey])
	if err != nil {
		return false, err
	}
	return intIsIn(c, v.comparisonValues), nil
}

func setCompareValidatorMessage(validator interfaces.Validator, setter func(ctx interfaces.ValidatorContextGetter) string) {
	catalog.ValidationCatalog().MessageStore().SetMessage(validator, func(ctx interfaces.ValidatorContextGetter) string {
		if compared, ok := ctx.GetContextData()[contextIsComparableTypesKey].(bool); ok && compared {
			return setter(ctx)
		}
		return fmt.Sprintf("Cannot compare because of incompatible comparative types.")
	})
}

func newCompareValidatorForValue(fieldName string, value interface{}, compareValues []int, validatorType interfaces.Validator) *compareValidator {
	return newCompareValidatorForContext(fieldName, validatorType, &valueCompare{
		getValue:            func(ctx interfaces.ValidatorContextGetter) interface{} { return value },
		comparisonValues:    compareValues,
		compareToContextKey: contextCompareToValueKey,
	})
}

func newCompareValidatorForValueAgainstOtherField(fieldName, otherFieldName string, compareValues []int, validatorType interfaces.Validator) *compareValidator {
	return newCompareValidatorForContext(fieldName, validatorType, &valueCompare{
		getValue:            func(ctx interfaces.ValidatorContextGetter) interface{} { return ctx.GetFieldValue(otherFieldName) },
		comparisonValues:    compareValues,
		compareToContextKey: contextCompareToValueKey,
	})
}

func newCompareValidatorForContext(fieldName string, validatorType interfaces.Validator, comparisons ...*valueCompare) *compareValidator {
	return &compareValidator{
		AllFieldValidators: &AllFieldValidators{
			fieldName: fieldName,
		},
		test: func(ctx *ValidatorContext) (result bool, err error) {
			ctx.AddContextData(ContextFieldValueKey, ctx.GetFieldValue(fieldName))

			for _, v := range comparisons {
				// add the compare to value to the context
				ctx.AddContextData(v.compareToContextKey, v.getValue(ctx))

				valid, err := v.evaluate(ctx)
				if err != nil {
					ctx.AddContextData(contextIsComparableTypesKey, false)
					return false, err
				}

				if !valid {
					return false, nil
				}
			}

			ctx.AddContextData(contextIsComparableTypesKey, true)
			return true, nil
		},
		validatorType: validatorType,
	}
}

func (v *compareValidator) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	ctx := v.AllFieldValidators.NewValidatorContext(thing, contextData)
	//nolint:errcheck // message store returns message if there is an error (based on context)
	if valid, _ := v.test(ctx); !valid {
		msg := messageStore.GetMessage(v.validatorType, ctx)
		return errors.NewValidationError(v.fieldName, msg)
	}
	return nil
}

func intIsIn(intValue int, intValues []int) bool {
	for _, v := range intValues {
		if intValue == v {
			return true
		}
	}
	return false
}

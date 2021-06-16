// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"reflect"

	"github.com/rbell/gospecexpress/errors"

	"github.com/rbell/gospecexpress/internal/reflectionhelpers"

	"github.com/rbell/gospecexpress/catalog"
	"github.com/rbell/gospecexpress/interfaces"
)

const (
	defaultContainsMessage = "%v should contain %v."
	contextContainsKey     = "Contains"
)

// Contains is a validator testing if a slice contains a value / reference
type Contains struct {
	*AllFieldValidators
	contains    interface{}
	containsVal reflect.Value
	fromContext interfaces.ValueFromContext
}

// NewContainsValidator returns an initialized contains validator
func NewContainsValidator(fieldName, alias string, value interface{}) interfaces.Validator {
	return &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		contains:    value,
		containsVal: reflect.ValueOf(value),
	}
}

// NewContainsValidatorFromContext returns an initialized contains validator
func NewContainsValidatorFromContext(fieldName, alias string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	return &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		fromContext: valueFromContext,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&Contains{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		//nolint:errcheck // ignore error
		val := ctx.GetContextData()[contextMaxLenKey]
		//nolint:errcheck // ignore error
		alias := ctx.GetContextData()[ContextFieldAliasKey].(string)
		return fmt.Sprintf(defaultContainsMessage, alias, val)
	})
}

// Validate asserts that the value is found in the slice
func (v *Contains) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if v.fromContext != nil {
		ctx := v.AllFieldValidators.NewValidatorContext(thing, contextData)
		v.contains = v.fromContext(ctx)
		v.containsVal = reflect.ValueOf(v.contains)
	}

	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.Kind() == reflect.Array || fv.Kind() == reflect.Slice || fv.Kind() == reflect.String {
			for i := 0; i < fv.Len(); i++ {
				item := fv.Index(i)
				//nolint:errcheck // ignore returned error - if the values cannot be compared the item does not match the value
				eq, _ := reflectionhelpers.Eq(item, v.containsVal)
				if eq {
					// we found it - return nil
					return nil
				}
			}
		}

		// we didn't find it, return error
		msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
			contextContainsKey: v.contains,
		}))
		return errors.NewValidationError(v.fieldName, msg)
	}

	return nil
}

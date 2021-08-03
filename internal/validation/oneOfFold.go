// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rbell/gospecexpress/catalog"
	"github.com/rbell/gospecexpress/errors"
	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/reflectionhelpers"
)

const (
	defaultOneOfFoldMessage = "%v does not match the required valid values %v (ignoring case)."
	oneOfFoldValuesKey      = "oneOfFoldValuesKey"
)

// OneOfFold defines validator enforcing a string must match one of a slice of strings, ignoring case
type OneOfFold struct {
	*AllFieldValidators
	values []string
}

// NewOneOfFold returns an initialized OneOf validator
func NewOneOfFold(fieldName, alias string, values []string) interfaces.Validator {
	return &OneOfFold{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		values: values,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&OneOfFold{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		//nolint:errcheck // ignore error
		values := ctx.GetContextData()[oneOfFoldValuesKey].([]interface{})
		//nolint:errcheck // ignore error
		alias := ctx.GetContextData()[ContextFieldAliasKey].(string)
		return fmt.Sprintf(defaultOneOfFoldMessage, alias, values)
	})
}

// Validate validates the field matches the regex
func (v *OneOfFold) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		// dereference if need be
		if fv.Kind() == reflect.Ptr {
			elemVal := fv.Elem()
			fv = &elemVal
		}

		// Short circuit if the field value is not a string since folding is only applicable to strings
		if fv.Kind() != reflect.String {
			return nil
		}

		fieldStringVal := fv.String()
		matches := false
		for _, val := range v.values {
			if strings.EqualFold(val, fieldStringVal) {
				matches = true
				break
			}
		}
		if !matches {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				oneOfFoldValuesKey: v.values,
			}))
			return errors.NewValidationError(v.fieldName, msg, v.shouldWarn)
		}
	}
	return nil
}

// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"reflect"

	"github.com/rbell/gospecexpress/catalog"
	"github.com/rbell/gospecexpress/errors"
	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/reflectionhelpers"
)

const (
	defaultOneOfMessage = "%v does not match the required valid values %v."
	oneOfValuesKey      = "oneOfValuesKey"
)

// OneOf defines validator enforcing a string must match a regex
type OneOf struct {
	*AllFieldValidators
	values []interface{}
}

// NewOneOf returns an initialized OneOf validator
func NewOneOf(fieldName, alias string, values []interface{}) interfaces.Validator {
	return &OneOf{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		values: values,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&OneOf{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		//nolint:errcheck // ignore error
		values := ctx.GetContextData()[oneOfValuesKey].([]interface{})
		//nolint:errcheck // ignore error
		alias := ctx.GetContextData()[ContextFieldAliasKey].(string)
		return fmt.Sprintf(defaultOneOfMessage, alias, values)
	})
}

// Validate validates the field matches the regex
func (v *OneOf) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		matches := false
		for _, val := range v.values {
			valval := reflect.ValueOf(val)
			if fv != nil {
				//nolint:errcheck // ignore eq error
				eq, _ := reflectionhelpers.Eq(*fv, valval)
				if eq {
					matches = true
					break
				}
			}
		}
		if !matches {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				oneOfValuesKey: v.values,
			}))
			return errors.NewValidationError(v.fieldName, msg, v.shouldWarn)
		}
	}
	return nil
}

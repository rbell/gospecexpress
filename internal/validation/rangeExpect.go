// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"reflect"

	"github.com/rbell/gospecexpress/errors"
	"github.com/rbell/gospecexpress/internal/errorhelpers"

	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/reflectionhelpers"
)

// RangeExpect allows a custom validation function to be applied over slice or array
type RangeExpect struct {
	*AllFieldValidators
	exp func(ctx interfaces.FieldValidatorContextGetter) error
}

// NewRangeExpect returns an initialized RangeExpect
func NewRangeExpect(fieldName, alias string, exp func(ctx interfaces.FieldValidatorContextGetter) error) interfaces.Validator {
	return &RangeExpect{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		exp: exp,
	}
}

// Validate validates the thing applying exp to each element
func (v *RangeExpect) Validate(thing interface{}, contextData map[string]interface{}, _ interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		// Only validate if array or slice
		if fv.Kind() == reflect.Array || fv.Kind() == reflect.Slice {
			var e error
			for i := 0; i < fv.Len(); i++ {
				item := fv.Index(i)
				ctx := v.AllFieldValidators.NewValidatorContext(item, contextData)
				err := v.exp(ctx)
				if err != nil {
					if ve, ok := err.(*errors.ValidatorError); ok {
						e = errorhelpers.JoinErrors(e, errors.NewValidationErrors(nil, map[string]*errors.ValidatorError{
							fmt.Sprintf("%v[%d]", v.fieldName, i): ve,
						}))
					}
				}
			}
			if e != nil && !reflect.ValueOf(e).IsNil() {
				return e
			}
		}
	}

	return nil
}

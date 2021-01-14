// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"reflect"

	"gitlab.com/rbell/gospecexpress/catalog"
	"gitlab.com/rbell/gospecexpress/errors"
	"gitlab.com/rbell/gospecexpress/interfaces"
	"gitlab.com/rbell/gospecexpress/internal/errorhelpers"
	"gitlab.com/rbell/gospecexpress/internal/reflectionhelpers"
)

// RangeValidate validates each element in an array or slice
type RangeValidate struct {
	*AllFieldValidators
	validationCatalog interfaces.Cataloger
}

// NewRangeValidate creates an initialized RangeValidate
func NewRangeValidate(fieldName, alias string) interfaces.Validator {
	return &Reference{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		validationCatalog: catalog.ValidationCatalog(),
	}
}

// Validate validates the reference
func (v *RangeValidate) Validate(thing interface{}, contextData map[string]interface{}, _ interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		// Only validate if array or slice
		if fv.Kind() == reflect.Array || fv.Kind() == reflect.Slice {
			var e error
			for i := 0; i < fv.Len(); i++ {
				item := fv.Index(i)
				err := v.validationCatalog.ValidateWithContext(item.Interface(), contextData)
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

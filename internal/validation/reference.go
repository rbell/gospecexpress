// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"github.com/rbell/gospecexpress/catalog"
	"github.com/rbell/gospecexpress/errors"
	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/reflectionhelpers"
)

// Reference validates a struct's reference to another type (i.e. Order.ShipAddress)
type Reference struct {
	*AllFieldValidators
	validationCatalog interfaces.Cataloger
}

// NewReferenceValidator returns an initialized Reference validator
func NewReferenceValidator(fieldName, alias string) interfaces.Validator {
	return &Reference{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		validationCatalog: catalog.ValidationCatalog(),
	}
}

// Validate validates the reference
func (v *Reference) Validate(thing interface{}, contextData map[string]interface{}, _ interfaces.MessageStorer) error {
	// get the value of the reference and validate it
	if val, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		err := v.validationCatalog.ValidateWithContext(val.Interface(), contextData)
		if err != nil {
			if ve, ok := err.(*errors.ValidatorError); ok {
				return errors.NewValidationErrors(nil, map[string]*errors.ValidatorError{
					v.fieldName: ve,
				})
			}
		}
	}

	return nil
}

// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"reflect"

	specExpressErrors "gitlab.com/rbell/gospecexpress/pkg/errors"
	"gitlab.com/rbell/gospecexpress/pkg/internal/errorhelpers"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

// Specification defines a base for specification
type Specification struct {
	forType    reflect.Type
	validators []interfaces.Validator
}

// ForType sets the type that the specification is to be applied to
func (s *Specification) ForType(forType interface{}) interfaces.QualifierBuilder {
	forValue := reflect.ValueOf(forType)
	s.forType = forValue.Type()
	s.validators = []interfaces.Validator{}
	return NewQualifierBuilder(&s.validators, forValue)
}

// GetForType returns the type that the specification is to be applied to
func (s *Specification) GetForType() reflect.Type {
	return s.forType
}

// Validate validates an instance of the type
func (s *Specification) Validate(thing interface{}, contextData map[string]interface{}) error {
	var specError *specExpressErrors.ValidatorError = nil
	for _, v := range s.validators {
		if err := v.Validate(thing, contextData, catalog.ValidationCatalog().MessageStore()); err != nil {
			specError = errorhelpers.JoinErrors(specError, err)
		}
	}

	if specError == nil || reflect.ValueOf(specError).IsNil() {
		return nil
	}
	return specError
}

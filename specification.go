// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gospecexpress

import (
	"reflect"
	"sync"

	"github.com/rbell/gospecexpress/internal/reflectionhelpers"

	specExpressErrors "github.com/rbell/gospecexpress/errors"
	"github.com/rbell/gospecexpress/internal/errorhelpers"

	"github.com/rbell/gospecexpress/catalog"

	"github.com/rbell/gospecexpress/interfaces"
)

// Specification defines a base for specification
type Specification struct {
	forType           reflect.Type
	fieldValidators   *sync.Map
	customExpressions []interfaces.ValidationExpression
}

// ForType sets the type that the specification is to be applied to
func (s *Specification) ForType(forType interface{}) interfaces.QualifierBuilder {
	forValue := reflect.ValueOf(forType)
	s.forType = forValue.Type()
	s.fieldValidators = &sync.Map{}
	return NewQualifierBuilder(s, forValue)
}

// GetForType returns the type that the specification is to be applied to
func (s *Specification) GetForType() reflect.Type {
	return s.forType
}

// Validate validates an instance of the type
func (s *Specification) Validate(thing interface{}, contextData map[string]interface{}) error {
	var specError *specExpressErrors.ValidatorError = nil

	// Validate customExpressions (expressions applied to structure as a whole)
	for _, exp := range s.customExpressions {
		if err := exp(thing, contextData); err != nil {
			specError = errorhelpers.JoinErrors(specError, err)
		}
	}

	// Validate field validators defined for the structure
	s.fieldValidators.Range(func(key, value interface{}) bool {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fieldName, _ := key.(string)
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fv, _ := value.(*fieldValidator)
		if fv.condition == nil || fv.condition(thing, contextData) {
			if fieldValue, ok := reflectionhelpers.GetFieldValue(thing, fieldName); ok {
				if fv.isOptional && fieldValue.IsZero() {
					// skip any validation since value in field is zero value and the field was optional
					return true
				}
				for _, v := range fv.validators {
					if err := v.Validate(thing, contextData, catalog.ValidationCatalog().MessageStore()); err != nil {
						specError = errorhelpers.JoinErrors(specError, err)
					}
				}
			}
		}
		return true
	})

	if specError == nil || reflect.ValueOf(specError).IsNil() {
		return nil
	}
	return specError
}

func addFieldValidator(fieldValidators *sync.Map, fieldName, alias string, validator interfaces.Validator) {
	var fv *fieldValidator
	if v, ok := fieldValidators.Load(fieldName); ok {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fv, _ = v.(*fieldValidator)
	} else {
		fv = &fieldValidator{
			isOptional: false,
			alias:      alias,
			validators: []interfaces.Validator{},
			mux:        &sync.Mutex{},
		}
	}
	fv.addValidator(validator)
	fieldValidators.Store(fieldName, fv)
}

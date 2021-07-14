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
	fieldExpressions  *sync.Map
	customExpressions []interfaces.ValidationExpression
}

// ForType sets the type that the specification is to be applied to
func (s *Specification) ForType(forType interface{}) interfaces.QualifierBuilder {
	forValue := reflect.ValueOf(forType)
	s.forType = forValue.Type()
	s.fieldExpressions = &sync.Map{}
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
		if verr, err := exp(thing, contextData); verr != nil || err != nil {
			if err != nil {
				return err
			}
			if verr != nil {
				specError = errorhelpers.JoinErrors(specError, err)
			}
		}
	}

	var processingError error

	// Validate field validators defined for the structure
	s.fieldExpressions.Range(func(key, value interface{}) bool {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fieldName, _ := key.(string)
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fexps, _ := value.([]*fieldExpression)
		for _, exp := range fexps {
			if exp.condition == nil || exp.condition(thing, contextData) {
				if fieldValue, ok := reflectionhelpers.GetFieldValue(thing, fieldName); ok {
					if exp.isOptional && fieldValue.IsZero() {
						// skip any validation since value in field is zero value and the field was optional
						return true
					}
					for _, v := range exp.validators {
						if err := v.Validate(thing, contextData, catalog.ValidationCatalog().MessageStore()); err != nil {
							if specErr, ok := specExpressErrors.IsValidatorError(err); ok {
								specError = errorhelpers.JoinErrors(specError, specErr)
							} else {
								processingError = err
								return false
							}
						}
					}
				}
			}
		}
		return true
	})

	if processingError != nil {
		return processingError
	}

	if specError == nil || reflect.ValueOf(specError).IsNil() {
		return nil
	}
	return specError
}

func addFieldExpression(fieldExpressions *sync.Map, fieldName, alias string, validator interfaces.Validator) {
	fe := &fieldExpression{
		isOptional: false,
		alias:      alias,
		validators: []interfaces.Validator{},
		mux:        &sync.Mutex{},
	}
	if v, ok := fieldExpressions.Load(fieldName); ok {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fexps, _ := v.([]*fieldExpression)
		fexps = append(fexps, fe)
		fieldExpressions.Store(fieldName, fexps)
	} else {
		fieldExpressions.Store(fieldName, []*fieldExpression{fe})
	}
	fe.addValidator(validator)
}

func addOptionalFieldExpression(fieldExpressions *sync.Map, fieldName, alias string) {
	fe := &fieldExpression{
		isOptional: true,
		alias:      alias,
		validators: []interfaces.Validator{},
		mux:        &sync.Mutex{},
	}
	if v, ok := fieldExpressions.Load(fieldName); ok {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fexps, _ := v.([]*fieldExpression)
		fexps = append(fexps, fe)
		fieldExpressions.Store(fieldName, fexps)
	} else {
		fieldExpressions.Store(fieldName, []*fieldExpression{fe})
	}
}

func addFieldValidator(fieldExpressions *sync.Map, fieldName, alias string, validator interfaces.Validator) {
	var fe *fieldExpression
	if v, ok := fieldExpressions.Load(fieldName); ok {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fexps, _ := v.([]*fieldExpression)
		fe = fexps[len(fexps)-1]
	} else {
		fe = &fieldExpression{
			isOptional: false,
			alias:      alias,
			validators: []interfaces.Validator{},
			mux:        &sync.Mutex{},
		}
		fieldExpressions.Store(fieldName, []*fieldExpression{fe})
	}
	fe.addValidator(validator)
}

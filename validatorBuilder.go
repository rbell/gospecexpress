// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gospecexpress

import (
	"reflect"
	"sync"

	"github.com/rbell/gospecexpress/internal/validation"

	"github.com/rbell/gospecexpress/interfaces"
)

type validatorBuilder struct {
	fieldName        string
	fieldAlias       string
	validators       *sync.Map
	forType          reflect.Value
	qualifierBuilder interfaces.QualifierBuilder
}

var _ interfaces.ValidatorBuilder = &validatorBuilder{}

// NewValidatorBuilder creates an initialized ValidatorBuilder
func NewValidatorBuilder(vals *sync.Map, forType reflect.Value, forField, alias string, builder interfaces.QualifierBuilder) interfaces.ValidatorBuilder {
	return &validatorBuilder{
		fieldName:        forField,
		fieldAlias:       alias,
		validators:       vals,
		forType:          forType,
		qualifierBuilder: builder,
	}
}

// Required indicates we want to start a new rule chain for a new required field
func (v *validatorBuilder) Required(fieldName string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	return v.qualifierBuilder.Required(fieldName, options...)
}

// Optional indicates we want to start a new rule chain for a field that is optional (not required).
func (v *validatorBuilder) Optional(fieldName string) interfaces.ValidatorBuilder {
	return v.qualifierBuilder.Optional(fieldName)
}

// If adds a condition as to when the rules for the rules we are building should be applied when validating
func (v *validatorBuilder) If(condition interfaces.FieldValidationCondition) interfaces.ValidatorBuilder {
	setCondition(v.validators, v.fieldName, condition)
	return v
}

func (v *validatorBuilder) OneOf(values []interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addFieldValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewOneOf(v.fieldName, v.fieldAlias, values), options...))
	return v
}

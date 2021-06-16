// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gospecexpress

import (
	"reflect"

	"github.com/rbell/gospecexpress/internal/reflectionhelpers"

	"github.com/rbell/gospecexpress/interfaces"

	"github.com/rbell/gospecexpress/internal/validation"
)

// NewQualifierBuilder creates an initialized QualifierBuilder
func NewQualifierBuilder(spec *Specification, forType reflect.Value) interfaces.QualifierBuilder {
	return &qualifierBuilder{
		spec:    spec,
		forType: forType,
	}
}

var _ interfaces.QualifierBuilder = &qualifierBuilder{}

// qualifierBuilder exists solely for the purpose of supporting code assistance enforcing functions available after `ForType`
type qualifierBuilder struct {
	spec    *Specification
	forType reflect.Value
}

// RequiredField indicates a field is required
func (b *qualifierBuilder) Required(fieldName string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	setOptional(b.spec.fieldValidators, fieldName, false)
	alias := reflectionhelpers.GetFieldAlias(b.forType, fieldName)
	addFieldValidator(b.spec.fieldValidators, fieldName, alias, ApplyValidatorOptions(validation.NewRequiredFieldValidator(fieldName, alias), options...))
	return NewValidatorBuilder(b.spec.fieldValidators, b.forType, fieldName, alias, b)
}

// Optional indicates a field is optional
func (b *qualifierBuilder) Optional(fieldName string) interfaces.ValidatorBuilder {
	setOptional(b.spec.fieldValidators, fieldName, true)
	alias := reflectionhelpers.GetFieldAlias(b.forType, fieldName)
	return NewValidatorBuilder(b.spec.fieldValidators, b.forType, alias, fieldName, b)
}

func (b *qualifierBuilder) Custom(exp interfaces.ValidationExpression) interfaces.QualifierBuilder {
	b.spec.customExpressions = append(b.spec.customExpressions, exp)
	return b
}

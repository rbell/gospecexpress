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

// ForScope sets the scope for the specification, allowing more than one specification be defined for a type.
// To validate using scope, use catalog.Validate(thing, catalog.WithScope("MyRegisteredScope"))
func (b *qualifierBuilder) ForScope(scope string) interfaces.QualifierBuilder {
	b.spec.scope = scope
	return b
}

// RequiredField indicates a field is required
func (b *qualifierBuilder) Required(fieldName string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	alias := reflectionhelpers.GetFieldAlias(b.forType, fieldName)
	addFieldExpression(b.spec.fieldExpressions, fieldName, alias, ApplyValidatorOptions(validation.NewRequiredFieldValidator(fieldName, alias), options...))
	return NewValidatorBuilder(b.spec.fieldExpressions, b.forType, fieldName, alias, b)
}

// Optional indicates a field is optional
func (b *qualifierBuilder) Optional(fieldName string) interfaces.ValidatorBuilder {
	alias := reflectionhelpers.GetFieldAlias(b.forType, fieldName)
	addOptionalFieldExpression(b.spec.fieldExpressions, fieldName, alias)
	return NewValidatorBuilder(b.spec.fieldExpressions, b.forType, alias, fieldName, b)
}

func (b *qualifierBuilder) Custom(exp interfaces.ValidationExpression) interfaces.QualifierBuilder {
	b.spec.customExpressions = append(b.spec.customExpressions, exp)
	return b
}

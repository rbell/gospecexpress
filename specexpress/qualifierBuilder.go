// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"reflect"
	"sync"

	"github.com/rbell/gospecexpress/internal/reflectionhelpers"

	"github.com/rbell/gospecexpress/interfaces"

	"github.com/rbell/gospecexpress/internal/validation"
)

// NewQualifierBuilder creates an initialized QualifierBuilder
func NewQualifierBuilder(vals *sync.Map, forType reflect.Value) interfaces.QualifierBuilder {
	return &qualifierBuilder{
		validators: vals,
		forType:    forType,
	}
}

var _ interfaces.QualifierBuilder = &qualifierBuilder{}

// qualifierBuilder exists solely for the purpose of supporting code assistance enforcing functions available after `ForType`
type qualifierBuilder struct {
	validators *sync.Map
	forType    reflect.Value
}

// RequiredField indicates a field is required
func (b *qualifierBuilder) Required(fieldName string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	setOptional(b.validators, fieldName, false)
	alias := reflectionhelpers.GetFieldAlias(b.forType, fieldName)
	addValidator(b.validators, fieldName, alias, ApplyValidatorOptions(validation.NewRequiredFieldValidator(fieldName, alias), options...))
	return NewValidatorBuilder(b.validators, b.forType, fieldName, alias, b)
}

// Optional indicates a field is optional
func (b *qualifierBuilder) Optional(fieldName string) interfaces.ValidatorBuilder {
	setOptional(b.validators, fieldName, true)
	alias := reflectionhelpers.GetFieldAlias(b.forType, fieldName)
	return NewValidatorBuilder(b.validators, b.forType, alias, fieldName, b)
}

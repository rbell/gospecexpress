// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"reflect"
	"sync"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// NewQualifierBuilder creates an initialized ValidatorBuilder
func NewQualifierBuilder(vals *sync.Map, forType reflect.Value) interfaces.QualifierBuilder {
	return &qualifierBuilder{
		validators: vals,
		forType:    forType,
	}
}

var _ interfaces.QualifierBuilder = &qualifierBuilder{}

type qualifierBuilder struct {
	validators *sync.Map
	forType    reflect.Value
}

// RequiredField indicates a field is required
func (b *qualifierBuilder) Required(fieldName string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	setOptional(b.validators, fieldName, false)
	addValidator(b.validators, fieldName, ApplyValidatorOptions(validation.NewRequiredFieldValidator(fieldName), options...))
	return NewValidatorBuilder(b.validators, b.forType, fieldName, b)
}

func (b *qualifierBuilder) Optional(fieldName string) interfaces.ValidatorBuilder {
	setOptional(b.validators, fieldName, true)
	return NewValidatorBuilder(b.validators, b.forType, fieldName, b)
}

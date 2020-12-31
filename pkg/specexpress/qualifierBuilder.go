// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// NewQualifierBuilder creates an initialized ValidatorBuilder
func NewQualifierBuilder(vals *[]interfaces.Validator, forType reflect.Value) interfaces.QualifierBuilder {
	return &qualifierBuilder{
		validators: vals,
		forType:    forType,
	}
}

var _ interfaces.QualifierBuilder = &qualifierBuilder{}

type qualifierBuilder struct {
	validators *[]interfaces.Validator
	forType    reflect.Value
}

// RequiredField indicates a field is required
func (b *qualifierBuilder) RequiredField(fieldName string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	vals := append(*b.validators, ApplyValidatorOptions(validation.NewRequiredFieldValidator(fieldName), options...))
	*b.validators = vals
	return NewValidatorBuilder(b.validators, b.forType, fieldName, b)
}

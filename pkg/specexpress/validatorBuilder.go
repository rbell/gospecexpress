// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

type validatorBuilder struct {
	fieldName        string
	validators       *[]interfaces.Validator
	forType          reflect.Value
	qualifierBuilder interfaces.QualifierBuilder
}

var _ interfaces.ValidatorBuilder = &validatorBuilder{}

// NewValidatorBuilder creates an initialized ValidatorBuilder
func NewValidatorBuilder(vals *[]interfaces.Validator, forType reflect.Value, forField string, builder interfaces.QualifierBuilder) interfaces.ValidatorBuilder {
	return &validatorBuilder{
		fieldName:        forField,
		validators:       vals,
		forType:          forType,
		qualifierBuilder: builder,
	}
}

// RequiredField indicates we want to start a new rule chain for a new required field
func (v *validatorBuilder) Required(fieldName string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	return v.qualifierBuilder.RequiredField(fieldName, options...)
}

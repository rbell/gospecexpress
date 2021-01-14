// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gospecexpress

import (
	"reflect"
	"sync"

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

func (v *validatorBuilder) Optional(fieldName string) interfaces.ValidatorBuilder {
	return v.qualifierBuilder.Optional(fieldName)
}

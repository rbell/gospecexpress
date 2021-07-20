// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"github.com/rbell/gospecexpress/errors"
	"github.com/rbell/gospecexpress/interfaces"
)

// Expectation is a validator wrapping a function used to validate a field
type Expectation struct {
	*AllFieldValidators
	exp interfaces.FieldValidationExpression
}

// NewExpectationValidator returns an initialized Expectation
func NewExpectationValidator(fieldName, alias string, exp interfaces.FieldValidationExpression) interfaces.Validator {
	return &Expectation{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		exp: exp,
	}
}

// Validate validates the thing ensuring the field specified is populated
func (e *Expectation) Validate(thing interface{}, contextData map[string]interface{}, _ interfaces.MessageStorer) error {
	ctx := e.AllFieldValidators.NewValidatorContext(thing, contextData)
	verr, err := e.exp(thing, ctx)
	if err != nil {
		return err
	} else if verr != nil {
		return errors.NewValidationError(e.AllFieldValidators.fieldName, verr.Error(), e.shouldWarn)
	}

	return nil
}

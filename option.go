// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gospecexpress

import (
	"github.com/rbell/gospecexpress/interfaces"
)

// ApplyValidatorOptions applies options to a validator, returning the optioned validator
func ApplyValidatorOptions(v interfaces.Validator, options ...interfaces.ValidatorOption) interfaces.Validator {
	for _, o := range options {
		o(v)
	}
	return v
}

// WithErrorMessageFormatter overrides the error message if the validator fails
func WithErrorMessageFormatter(msgFormatter interfaces.MessageFormatter) func(validator interfaces.Validator) {
	return func(validator interfaces.Validator) {
		if overrider, ok := validator.(interfaces.MessageOverrider); ok {
			overrider.SetOverrideErrorMessage(msgFormatter)
		}
	}
}

// WithErrorMessage overrides the error message if the validator fails
func WithErrorMessage(msg string) func(validator interfaces.Validator) {
	return WithErrorMessageFormatter(func(ctx interfaces.FieldValidatorContextGetter) string {
		return msg
	})
}

// WithWarning indicates that any resulting failed validation should result in a warning
func WithWarning() func(validator interfaces.Validator) {
	return func(validator interfaces.Validator) {
		if warningSetter, ok := validator.(interfaces.WarningSetter); ok {
			warningSetter.ValidateAsWarning()
		}

	}
}

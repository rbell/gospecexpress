// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/validation"
)

// Expect provides a way to express a function that should be used to validate a field
func (v *validatorBuilder) Expect(validatorFunc func(valueFromContext interfaces.ValidatorContextGetter) error, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewExpectationValidator(v.fieldName, v.fieldAlias, validatorFunc), options...))
	return v
}

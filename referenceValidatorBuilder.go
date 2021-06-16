// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gospecexpress

import (
	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/validation"
)

// ValidateReference forces validating the value against the catalog (i.e. validate the address struct referenced by customer in the customer's Address field)
func (v *validatorBuilder) ValidateReference(options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addFieldValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewReferenceValidator(v.fieldName, v.fieldAlias), options...))
	return v
}

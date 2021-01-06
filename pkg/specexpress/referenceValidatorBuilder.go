// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// ValidateReference forces validating the value against the catalog (i.e. validate the address struct referenced by customer in the customer's Address field)
func (v *validatorBuilder) ValidateReference(options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.NewReferenceValidator(v.fieldName), options...))
	return v
}

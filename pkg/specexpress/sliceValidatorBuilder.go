// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file..

package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// LengthEquals indicates a lenth equal to rule should be applied to the field
func (v *validatorBuilder) LengthEquals(length int) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, validation.NewLengthEqualsValidator(v.fieldName, length))
	return v
}

// MaxLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MaxLength(length int) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, validation.NewMaxLengthValidator(v.fieldName, length))
	return v
}

// MinLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MinLength(length int) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, validation.NewMinLengthValidator(v.fieldName, length))
	return v
}

// Contains validates the slice contains some thing
func (v *validatorBuilder) Contains(thing interface{}) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, validation.NewContainsValidator(v.fieldName, thing))
	return v
}

// Contains validates the slice contains some thing
func (v *validatorBuilder) ContainsValueFromContext(fromContext interfaces.ValueFromContext) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, validation.NewContainsValidatorFromContext(v.fieldName, fromContext))
	return v
}

// RangeValidate validates each element in the array or slice against the catalog
func (v *validatorBuilder) RangeValidate() interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, validation.NewRangeValidate(v.fieldName))
	return v
}

// RangeExpect allows a custom validation function to be applied over a slice or array
func (v *validatorBuilder) RangeExpect(validator func(validationCtx interfaces.ValidatorContextGetter) error) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, validation.NewRangeExpect(v.fieldName, validator))
	return v
}

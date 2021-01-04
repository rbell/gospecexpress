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
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.NewLengthEqualsValidator(v.fieldName, length))
	*v.validators = vals
	return v
}

// MaxLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MaxLength(length int) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.NewMaxLengthValidator(v.fieldName, length))
	*v.validators = vals
	return v
}

// MinLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MinLength(length int) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.NewMinLengthValidator(v.fieldName, length))
	*v.validators = vals
	return v
}

// Contains validates the slice contains some thing
func (v *validatorBuilder) Contains(thing interface{}) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.NewContainsValidator(v.fieldName, thing))
	*v.validators = vals
	return v
}

// Contains validates the slice contains some thing
func (v *validatorBuilder) ContainsValueFromContext(fromContext interfaces.ValueFromContext) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.NewContainsValidatorFromContext(v.fieldName, fromContext))
	*v.validators = vals
	return v
}

// RangeValidate validates each element in the array or slice against the catalog
func (v *validatorBuilder) RangeValidate() interfaces.ValidatorBuilder {
	vals := append(*v.validators, validation.NewRangeValidate(v.fieldName))
	*v.validators = vals
	return v
}

// RangeExpect allows a custom validation function to be applied over a slice or array
func (v *validatorBuilder) RangeExpect(validator func(validationCtx interfaces.ValidatorContextGetter) error) interfaces.ValidatorBuilder {
	vals := append(*v.validators, validation.NewRangeExpect(v.fieldName, validator))
	*v.validators = vals
	return v
}

// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file..

package specexpress

import (
	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/validation"
)

// LengthEquals indicates a lenth equal to rule should be applied to the field
func (v *validatorBuilder) LengthEquals(length int, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewLengthEqualsValidator(v.fieldName, v.fieldAlias, length), options...))
	return v
}

// MaxLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MaxLength(length int, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewMaxLengthValidator(v.fieldName, v.fieldAlias, length), options...))
	return v
}

// MinLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MinLength(length int, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewMinLengthValidator(v.fieldName, v.fieldAlias, length), options...))
	return v
}

// Contains validates the slice contains some thing
func (v *validatorBuilder) Contains(thing interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewContainsValidator(v.fieldName, v.fieldAlias, thing), options...))
	return v
}

// Contains validates the slice contains some thing
func (v *validatorBuilder) ContainsValueFromContext(fromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewContainsValidatorFromContext(v.fieldName, v.fieldAlias, fromContext), options...))
	return v
}

// RangeValidate validates each element in the array or slice against the catalog
func (v *validatorBuilder) RangeValidate(options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewRangeValidate(v.fieldName, v.fieldAlias), options...))
	return v
}

// RangeExpect allows a custom validation function to be applied over a slice or array
func (v *validatorBuilder) RangeExpect(validator func(validationCtx interfaces.ValidatorContextGetter) error, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewRangeExpect(v.fieldName, v.fieldAlias, validator), options...))
	return v
}

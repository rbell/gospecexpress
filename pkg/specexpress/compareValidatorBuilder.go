// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// Between ensures the value in the field is between the lower and upper parameters
func (v *validatorBuilder) Between(lower, upper interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.BetweenValues(v.fieldName, lower, upper), options...))
	return v
}

// Between ensures the value in the field is between values stored in the lowerField and upperField parameters
func (v *validatorBuilder) BetweenOtherFields(lowerField, upperField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.BetweenOtherFieldValues(v.fieldName, lowerField, upperField), options...))
	return v
}

// Between ensures the value in the field is between values stored in the validation context
func (v *validatorBuilder) BetweenValuesFromContext(lowerGetter, upperGetter interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.BetweenValuesFromContext(v.fieldName, lowerGetter, upperGetter), options...))
	return v
}

// LessThan indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThan(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.LessThanValue(v.fieldName, value), options...))
	return v
}

// LessThanOtherField indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThanOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.LessThanFieldValue(v.fieldName, otherField), options...))
	return v
}

// LessThanValueFromContext indicates a less than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) LessThanValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.LessThanValueFromContext(v.fieldName, valueFromContext), options...))
	return v
}

// LessThanOrEqualTo indicates a less than or equal to rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.LessThanOrEqualToValue(v.fieldName, value), options...))
	return v
}

// LessThanOrEqualToOtherField indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.LessThanOrEqualToFieldValue(v.fieldName, otherField), options...))
	return v
}

// LessThanValueFromContext indicates a less than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) LessThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.LessThanOrEqualToValueFromContext(v.fieldName, valueFromContext), options...))
	return v
}

// GreaterThan indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThan(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.GreaterThanValue(v.fieldName, value), options...))
	return v
}

// GreaterThanOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.GreaterThanFieldValue(v.fieldName, otherField), options...))
	return v
}

// GreaterThanValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) GreaterThanValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.GreaterThanValueFromContext(v.fieldName, valueFromContext), options...))
	return v
}

// GreaterThanOrEqualToOrEqualTo indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOrEqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.GreaterThanOrEqualToValue(v.fieldName, value), options...))
	return v
}

// GreaterThanOrEqualToOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOrEqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.GreaterThanOrEqualToFieldValue(v.fieldName, otherField), options...))
	return v
}

// GreaterThanOrEqualToValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) GreaterThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.GreaterThanOrEqualToValueFromContext(v.fieldName, valueFromContext), options...))
	return v
}

// EqualTo indicates a Greater than rule should be applied to field
func (v *validatorBuilder) EqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.EqualToValue(v.fieldName, value), options...))
	return v
}

// EqualToOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) EqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.EqualToFieldValue(v.fieldName, otherField), options...))
	return v
}

// EqualToValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) EqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, ApplyValidatorOptions(validation.EqualToValueFromContext(v.fieldName, valueFromContext), options...))
	return v
}

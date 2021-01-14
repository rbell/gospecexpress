// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/validation"
)

// Between ensures the value in the field is between the lower and upper parameters
func (v *validatorBuilder) Between(lower, upper interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.BetweenValues(v.fieldName, v.fieldAlias, lower, upper), options...))
	return v
}

// Between ensures the value in the field is between values stored in the lowerField and upperField parameters
func (v *validatorBuilder) BetweenOtherFields(lowerField, upperField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.BetweenOtherFieldValues(v.fieldName, v.fieldAlias, lowerField, upperField), options...))
	return v
}

// Between ensures the value in the field is between values stored in the validation context
func (v *validatorBuilder) BetweenValuesFromContext(lowerGetter, upperGetter interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.BetweenValuesFromContext(v.fieldName, v.fieldAlias, lowerGetter, upperGetter), options...))
	return v
}

// LessThan indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThan(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.LessThanValue(v.fieldName, v.fieldAlias, value), options...))
	return v
}

// LessThanOtherField indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThanOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.LessThanFieldValue(v.fieldName, v.fieldAlias, otherField), options...))
	return v
}

// LessThanValueFromContext indicates a less than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) LessThanValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.LessThanValueFromContext(v.fieldName, v.fieldAlias, valueFromContext), options...))
	return v
}

// LessThanOrEqualTo indicates a less than or equal to rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.LessThanOrEqualToValue(v.fieldName, v.fieldAlias, value), options...))
	return v
}

// LessThanOrEqualToOtherField indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.LessThanOrEqualToFieldValue(v.fieldName, v.fieldAlias, otherField), options...))
	return v
}

// LessThanValueFromContext indicates a less than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) LessThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.LessThanOrEqualToValueFromContext(v.fieldName, v.fieldAlias, valueFromContext), options...))
	return v
}

// GreaterThan indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThan(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.GreaterThanValue(v.fieldName, v.fieldAlias, value), options...))
	return v
}

// GreaterThanOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.GreaterThanFieldValue(v.fieldName, v.fieldAlias, otherField), options...))
	return v
}

// GreaterThanValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) GreaterThanValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.GreaterThanValueFromContext(v.fieldName, v.fieldAlias, valueFromContext), options...))
	return v
}

// GreaterThanOrEqualToOrEqualTo indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOrEqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.GreaterThanOrEqualToValue(v.fieldName, v.fieldAlias, value), options...))
	return v
}

// GreaterThanOrEqualToOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOrEqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.GreaterThanOrEqualToFieldValue(v.fieldName, v.fieldAlias, otherField), options...))
	return v
}

// GreaterThanOrEqualToValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) GreaterThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.GreaterThanOrEqualToValueFromContext(v.fieldName, v.fieldAlias, valueFromContext), options...))
	return v
}

// EqualTo indicates a Greater than rule should be applied to field
func (v *validatorBuilder) EqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.EqualToValue(v.fieldName, v.fieldAlias, value), options...))
	return v
}

// EqualToOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) EqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.EqualToFieldValue(v.fieldName, v.fieldAlias, otherField), options...))
	return v
}

// EqualToValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) EqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.EqualToValueFromContext(v.fieldName, v.fieldAlias, valueFromContext), options...))
	return v
}

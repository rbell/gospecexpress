// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// LessThan indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThan(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.LessThanValue(v.fieldName, value), options...))
	*v.validators = vals
	return v
}

// LessThanOtherField indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThanOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.LessThanFieldValue(v.fieldName, otherField), options...))
	*v.validators = vals
	return v
}

// LessThanValueFromContext indicates a less than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) LessThanValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.LessThanValueFromContext(v.fieldName, valueFromContext), options...))
	*v.validators = vals
	return v
}

// LessThanOrEqualTo indicates a less than or equal to rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.LessThanOrEqualToValue(v.fieldName, value), options...))
	*v.validators = vals
	return v
}

// LessThanOrEqualToOtherField indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.LessThanOrEqualToFieldValue(v.fieldName, otherField), options...))
	*v.validators = vals
	return v
}

// LessThanValueFromContext indicates a less than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) LessThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.LessThanOrEqualToValueFromContext(v.fieldName, valueFromContext), options...))
	*v.validators = vals
	return v
}

// GreaterThan indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThan(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.GreaterThanValue(v.fieldName, value), options...))
	*v.validators = vals
	return v
}

// GreaterThanOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.GreaterThanFieldValue(v.fieldName, otherField), options...))
	*v.validators = vals
	return v
}

// GreaterThanValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) GreaterThanValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.GreaterThanValueFromContext(v.fieldName, valueFromContext), options...))
	*v.validators = vals
	return v
}

// GreaterThanOrEqualToOrEqualTo indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOrEqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.GreaterThanOrEqualToValue(v.fieldName, value), options...))
	*v.validators = vals
	return v
}

// GreaterThanOrEqualToOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOrEqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.GreaterThanOrEqualToFieldValue(v.fieldName, otherField), options...))
	*v.validators = vals
	return v
}

// GreaterThanOrEqualToValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) GreaterThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.GreaterThanOrEqualToValueFromContext(v.fieldName, valueFromContext), options...))
	*v.validators = vals
	return v
}

// EqualTo indicates a Greater than rule should be applied to field
func (v *validatorBuilder) EqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.EqualToValue(v.fieldName, value), options...))
	*v.validators = vals
	return v
}

// EqualToOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) EqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.EqualToFieldValue(v.fieldName, otherField), options...))
	*v.validators = vals
	return v
}

// EqualToValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) EqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, ApplyValidatorOptions(validation.EqualToValueFromContext(v.fieldName, valueFromContext), options...))
	*v.validators = vals
	return v
}

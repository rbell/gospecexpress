package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// LessThan indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThan(value interface{}) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.LessThanValue(v.fieldName, value))
	*v.validators = vals
	return v
}

// LessThanOtherField indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThanOtherField(otherField string) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.LessThanFieldValue(v.fieldName, otherField))
	*v.validators = vals
	return v
}

// LessThanValueFromContext indicates a less than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) LessThanValueFromContext(valueFromContext interfaces.ValueFromContext) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.LessThanValueFromContext(v.fieldName, valueFromContext))
	*v.validators = vals
	return v
}

// LessThanOrEqualTo indicates a less than or equal to rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualTo(value interface{}) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.LessThanOrEqualToValue(v.fieldName, value))
	*v.validators = vals
	return v
}

// LessThanOrEqualToOtherField indicates a less than rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualToOtherField(otherField string) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.LessThanOrEqualToFieldValue(v.fieldName, otherField))
	*v.validators = vals
	return v
}

// LessThanValueFromContext indicates a less than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) LessThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.LessThanOrEqualToValueFromContext(v.fieldName, valueFromContext))
	*v.validators = vals
	return v
}

// GreaterThan indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThan(value interface{}) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.GreaterThanValue(v.fieldName, value))
	*v.validators = vals
	return v
}

// GreaterThanOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOtherField(otherField string) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.GreaterThanFieldValue(v.fieldName, otherField))
	*v.validators = vals
	return v
}

// GreaterThanValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) GreaterThanValueFromContext(valueFromContext interfaces.ValueFromContext) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.GreaterThanValueFromContext(v.fieldName, valueFromContext))
	*v.validators = vals
	return v
}

// GreaterThanOrEqualToOrEqualTo indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOrEqualTo(value interface{}) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.GreaterThanOrEqualToValue(v.fieldName, value))
	*v.validators = vals
	return v
}

// GreaterThanOrEqualToOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) GreaterThanOrEqualToOtherField(otherField string) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.GreaterThanOrEqualToFieldValue(v.fieldName, otherField))
	*v.validators = vals
	return v
}

// GreaterThanOrEqualToValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) GreaterThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.GreaterThanOrEqualToValueFromContext(v.fieldName, valueFromContext))
	*v.validators = vals
	return v
}

// EqualTo indicates a Greater than rule should be applied to field
func (v *validatorBuilder) EqualTo(value interface{}) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.EqualToValue(v.fieldName, value))
	*v.validators = vals
	return v
}

// EqualToOtherField indicates a Greater than rule should be applied to field
func (v *validatorBuilder) EqualToOtherField(otherField string) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.EqualToFieldValue(v.fieldName, otherField))
	*v.validators = vals
	return v
}

// EqualToValueFromContext indicates a Greater than rule should be applied to field comparing it to value from context
func (v *validatorBuilder) EqualToValueFromContext(valueFromContext interfaces.ValueFromContext) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.EqualToValueFromContext(v.fieldName, valueFromContext))
	*v.validators = vals
	return v
}
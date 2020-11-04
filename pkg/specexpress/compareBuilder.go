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

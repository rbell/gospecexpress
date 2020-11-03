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

// LessThanOrEqualTo indicates a less than or equal to rule should be applied to field
func (v *validatorBuilder) LessThanOrEqualTo(value interface{}) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.LessThanOrEqualToValue(v.fieldName, value))
	*v.validators = vals
	return v
}

package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// LessThan indicates a LessThan rule should be applied to field
func (v *validatorBuilder) LessThan(value interface{}) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.LessThanValue(v.fieldName, value))
	*v.validators = vals
	return v
}

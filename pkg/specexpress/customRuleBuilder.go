package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// Expect provides a way to express a function that should be used to validate a field
func (v *validatorBuilder) Expect(validatorFunc func(valueFromContext interfaces.ValidatorContextGetter) error) interfaces.ValidatorBuilder {
	vals := append(*v.validators, validation.NewExpectationValidator(v.fieldName, validatorFunc))
	*v.validators = vals
	return v
}

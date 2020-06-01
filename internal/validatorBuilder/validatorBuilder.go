package validatorBuilder

import (
	"gitlab.com/govalidate/internal/validators"
	"gitlab.com/govalidate/pkg/interfaces"
	"reflect"
)

func NewValidatorBuilder(validators *[]interfaces.Validator, forType reflect.Value) interfaces.ValidatorBuilder {
	return &builder{
		validators: validators,
		forType: forType,
	}
}

var _ interfaces.ValidatorBuilder = &builder{}

type builder struct {
	validators *[]interfaces.Validator
	forType reflect.Value
}

func (b *builder) RequiredField(fieldName string) interfaces.ValidatorBuilder {
	validators := append(*b.validators, validators.NewRequiredFieldValidator(fieldName))
	*b.validators = validators
	return b
}

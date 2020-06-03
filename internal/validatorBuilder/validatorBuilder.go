package validatorbuilder

import (
	"reflect"

	"gitlab.com/govalidate/internal/validators"
	"gitlab.com/govalidate/pkg/interfaces"
)

// NewValidatorBuilder creates an initialized ValidatorBuilder
func NewValidatorBuilder(vals *[]interfaces.Validator, forType reflect.Value) interfaces.ValidatorBuilder {
	return &builder{
		validators: vals,
		forType:    forType,
	}
}

var _ interfaces.ValidatorBuilder = &builder{}

type builder struct {
	validators *[]interfaces.Validator
	forType    reflect.Value
}

// RequiredField indicates a field is required
func (b *builder) RequiredField(fieldName string) interfaces.ValidatorBuilder {
	vals := append(*b.validators, validators.NewRequiredFieldValidator(fieldName))
	*b.validators = vals
	return b
}

package specexpress

import (
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// NewQualifierBuilder creates an initialized ValidatorBuilder
func NewQualifierBuilder(vals *[]interfaces.Validator, forType reflect.Value) interfaces.QualifierBuilder {
	return &qualifierBuilder{
		validators: vals,
		forType:    forType,
	}
}

var _ interfaces.QualifierBuilder = &qualifierBuilder{}

type qualifierBuilder struct {
	validators *[]interfaces.Validator
	forType    reflect.Value
}

// RequiredField indicates a field is required
func (b *qualifierBuilder) RequiredField(fieldName string) interfaces.ValidatorBuilder {
	vals := append(*b.validators, validation.NewRequiredFieldValidator(fieldName))
	*b.validators = vals
	return NewValidatorBuilder(b.validators, b.forType, fieldName, b)
}

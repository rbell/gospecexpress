package builders

import (
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

type validatorBuilder struct {
	fieldName        string
	validators       *[]interfaces.Validator
	forType          reflect.Value
	qualifierBuilder interfaces.QualifierBuilder
}

// NewValidatorBuilder creates an initialized ValidatorBuilder
func NewValidatorBuilder(vals *[]interfaces.Validator, forType reflect.Value, forField string, builder interfaces.QualifierBuilder) interfaces.ValidatorBuilder {
	return &validatorBuilder{
		fieldName:        forField,
		validators:       vals,
		forType:          forType,
		qualifierBuilder: builder,
	}
}

// RequiredField indicates we want to start a new rule chain for a new required field
func (v *validatorBuilder) RequiredField(fieldName string) interfaces.ValidatorBuilder {
	return v.qualifierBuilder.RequiredField(fieldName)
}

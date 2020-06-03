package specexpress

import (
	"reflect"

	"gitlab.com/govalidate/internal/validatorbuilder"
	"gitlab.com/govalidate/pkg/errors"
	"gitlab.com/govalidate/pkg/interfaces"
)

// Specification defines a base for specification
type Specification struct {
	forType    reflect.Type
	validators []interfaces.Validator
}

// ForType sets the type that the specification is to be applied to
func (s *Specification) ForType(forType interface{}) interfaces.ValidatorBuilder {
	forValue := reflect.ValueOf(forType)
	s.forType = forValue.Type()
	s.validators = []interfaces.Validator{}
	return validatorbuilder.NewValidatorBuilder(&s.validators, forValue)
}

// GetForType returns the type that the specification is to be applied to
func (s *Specification) GetForType() reflect.Type {
	return s.forType
}

// Validate validates an instance of the type
func (s *Specification) Validate(thing interface{}) error {
	var specError *errors.ValidationError = nil
	for _, v := range s.validators {
		if err := v.Validate(thing); err != nil {
			specError = errors.JoinErrors(specError, err)
		}
	}

	if specError == nil || reflect.ValueOf(specError).IsNil() {
		return nil
	}
	return specError
}

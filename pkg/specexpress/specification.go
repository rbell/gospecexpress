package specexpress

import (
	"gitlab.com/govalidate/internal/validatorBuilder"
	"gitlab.com/govalidate/pkg/interfaces"
	"reflect"
)

// Specification defines a base for specification
type Specification struct {
	forType reflect.Type
	validators []interfaces.Validator
}

// ForType sets the type that the specification is to be applied to
func (s *Specification) ForType(forType interface{}) interfaces.ValidatorBuilder {
	forValue := reflect.ValueOf(forType)
	s.forType = forValue.Type()
	s.validators = []interfaces.Validator{}
	return validatorBuilder.NewValidatorBuilder(&s.validators, forValue)
}

// GetForType returns the type that the specification is to be applied to
func (s *Specification) GetForType() reflect.Type {
	return s.forType
}

// Validate validates an instance of the type
func (s *Specification) Validate(thing interface{}) bool {
	for _, v := range s.validators {
		if !v.Validate(thing) {
			return false
		}
	}
	return true
}


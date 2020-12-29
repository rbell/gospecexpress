package specexpress

import (
	"errors"
	"reflect"

	specExpressErrors "gitlab.com/rbell/gospecexpress/pkg/errors"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

// Specification defines a base for specification
type Specification struct {
	forType    reflect.Type
	validators []interfaces.Validator
}

// ForType sets the type that the specification is to be applied to
func (s *Specification) ForType(forType interface{}) interfaces.QualifierBuilder {
	forValue := reflect.ValueOf(forType)
	s.forType = forValue.Type()
	s.validators = []interfaces.Validator{}
	return NewQualifierBuilder(&s.validators, forValue)
}

// GetForType returns the type that the specification is to be applied to
func (s *Specification) GetForType() reflect.Type {
	return s.forType
}

// Validate validates an instance of the type
func (s *Specification) Validate(thing interface{}, contextData map[string]interface{}) error {
	var specError *specExpressErrors.ValidatorError = nil
	for _, v := range s.validators {
		if err := v.Validate(thing, contextData, catalog.ValidationCatalog().MessageStore()); err != nil {
			specError = joinErrors(specError, err)
		}
	}

	if specError == nil || reflect.ValueOf(specError).IsNil() {
		return nil
	}
	return specError
}

func joinErrors(e1, e2 error) *specExpressErrors.ValidatorError {
	var e *specExpressErrors.ValidatorError
	if (e1 == nil || reflect.ValueOf(e1).IsNil()) && e2 != nil {
		if errors.As(e2, &e) {
			return e2.(*specExpressErrors.ValidatorError)
		}
		return specExpressErrors.NewValidationError("", e2.Error())

	}

	var ve *specExpressErrors.ValidatorError
	if errors.As(e1, &e) {
		//nolint:errcheck // above line infers its castable
		ve = e1.(*specExpressErrors.ValidatorError)
	} else {
		ve = specExpressErrors.NewValidationError("", e1.Error())
	}

	if errors.As(e2, &e) {
		errMap := ve.GetErrorMap()
		for key, msg := range e2.(*specExpressErrors.ValidatorError).GetFlatErrorMap() {
			addMsgs(errMap, key, msg...)
		}
		childErrs := ve.GetChildErrors()
		for key, ve := range e2.(*specExpressErrors.ValidatorError).GetChildErrors() {
			childErrs[key] = ve
		}
	}

	return ve
}

func addMsgs(errMap map[string][]string, context string, msg ...string) {
	if _, ok := errMap[context]; !ok {
		errMap[context] = []string{}
	}
	errMap[context] = append(errMap[context], msg...)
}

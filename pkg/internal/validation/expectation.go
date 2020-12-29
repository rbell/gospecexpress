package validation

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

// Expectation is a validator wrapping a function used to validate a field
type Expectation struct {
	*AllFieldValidators
	exp func(ctx interfaces.ValidatorContextGetter) error
}

// NewExpectationValidator returns an initialized Expectation
func NewExpectationValidator(fieldName string, exp func(ctx interfaces.ValidatorContextGetter) error) interfaces.Validator {
	return &Expectation{
		AllFieldValidators: &AllFieldValidators{
			fieldName: fieldName,
		},
		exp: exp,
	}
}

// Validate validates the thing ensuring the field specified is populated
func (e *Expectation) Validate(thing interface{}, contextData map[string]interface{}, _ interfaces.MessageStorer) error {
	ctx := e.AllFieldValidators.NewValidatorContext(thing, contextData)
	err := e.exp(ctx)
	if err != nil {
		return NewValidationError(e.AllFieldValidators.fieldName, err.Error())
	}

	return nil
}

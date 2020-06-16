package validators

import (
	"fmt"

	"gitlab.com/govalidate/internal/reflectionhelpers"
	"gitlab.com/govalidate/pkg/errors"
	"gitlab.com/govalidate/pkg/interfaces"
)

// MaxLength defines a validator testing the length of a field
type MaxLength struct {
	fieldName string
	maxLen    int
}

// NewMaxLengthValidator creates an initialized MaxLengthValidator
func NewMaxLengthValidator(fieldName string, maxLen int) interfaces.Validator {
	return &MaxLength{
		fieldName: fieldName,
		maxLen:    maxLen,
	}
}

// Validate validates the thing ensureing the field specified is populated
func (v *MaxLength) Validate(thing interface{}) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.Len() > v.maxLen {
			// TODO: Get message from a msg repository of some sorts
			return errors.NewValidationError(v.fieldName, fmt.Sprintf("%v should not have a length greater than %v.", v.fieldName, v.maxLen))
		}
	}

	return nil
}

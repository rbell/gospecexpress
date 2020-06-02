package validators

import (
	"fmt"
	"gitlab.com/govalidate/internal/reflectionHelpers"
	"gitlab.com/govalidate/pkg/errors"
	"gitlab.com/govalidate/pkg/interfaces"
)

// Required Field defines a validator requiring a field value be populated.
type RequiredField struct {
	fieldName string
}

func NewRequiredFieldValidator(fieldName string) interfaces.Validator {
	return &RequiredField{
		fieldName: fieldName,
	}
}

func (v *RequiredField) Validate(thing interface{}) error {
	if fv, ok := reflectionHelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.IsZero() {
			// TODO: Get message from a msg repository of some sorts
			return errors.NewValidationError(v.fieldName, fmt.Sprintf("%v is required.", v.fieldName))
		}
	}

	return nil
}

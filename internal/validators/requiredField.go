package validators

import (
	"gitlab.com/govalidate/internal/reflectionHelpers"
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

func (v *RequiredField) Validate(thing interface{}) bool {
	if fv, ok := reflectionHelpers.GetFieldValue(thing, v.fieldName); ok {
		return !fv.IsZero()
	}

	return false
}

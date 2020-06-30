package validators

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/specificationcatalog"

	"gitlab.com/rbell/gospecexpress/internal/reflectionhelpers"

	"gitlab.com/rbell/gospecexpress/pkg/errors"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultRequiredFieldMessage = "%v is required."

// RequiredField defines a validator requiring a field value be populated.
type RequiredField struct {
	*AllFieldValidators
}

// NewRequiredFieldValidator creates an initialized RequiredFieldValidator
func NewRequiredFieldValidator(fieldName string) interfaces.Validator {
	return &RequiredField{
		AllFieldValidators: &AllFieldValidators{
			FieldName: fieldName,
		},
	}
}

func init() {
	specificationcatalog.Catalog().MessageStore().SetMessage(&RequiredField{}, func(ctx *errors.ErrorMessageContext) string {
		return fmt.Sprintf(defaultRequiredFieldMessage, ctx.ContextData[0].(string))
	})
}

// Validate validates the thing ensureing the field specified is populated
func (v *RequiredField) Validate(thing interface{}) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.FieldName); ok {
		if fv.IsZero() {
			msg := specificationcatalog.Catalog().MessageStore().GetMessage(v, v.AllFieldValidators.ErrorMessageContext(thing))
			return errors.NewValidationError(v.FieldName, msg)
		}
	}

	return nil
}

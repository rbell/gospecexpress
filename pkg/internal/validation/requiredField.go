package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"
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
	catalog.ValidationCatalog().MessageStore().SetMessage(&RequiredField{}, func(ctx interfaces.ValidatorContextGetter) string {
		return fmt.Sprintf(defaultRequiredFieldMessage, ctx.GetContextData()[ContextFieldNameKey].(string))
	})
}

// Validate validates the thing ensureing the field specified is populated
func (v *RequiredField) Validate(thing interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.FieldName); ok {
		if fv.IsZero() {
			msg := catalog.ValidationCatalog().MessageStore().GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, nil))
			return NewValidationError(v.FieldName, msg)
		}
	}

	return nil
}

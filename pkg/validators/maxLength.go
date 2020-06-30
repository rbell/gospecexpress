package validators

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/specificationcatalog"

	"gitlab.com/rbell/gospecexpress/internal/reflectionhelpers"
	"gitlab.com/rbell/gospecexpress/pkg/errors"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultMaxLengthMessage = "%v should not have a length greater than %v."

// MaxLength defines a validator testing the length of a field
type MaxLength struct {
	*AllFieldValidators
	maxLen int
}

// NewMaxLengthValidator creates an initialized MaxLengthValidator
func NewMaxLengthValidator(fieldName string, maxLen int) interfaces.Validator {
	return &MaxLength{
		AllFieldValidators: &AllFieldValidators{
			FieldName: fieldName,
		},
		maxLen: maxLen,
	}
}

func init() {
	specificationcatalog.Catalog().MessageStore().SetMessage(&MaxLength{}, func(ctx *errors.ErrorMessageContext) string {
		fieldValue := ctx.GetFieldValue(ctx.ContextData[0].(string))
		//nolint:errcheck // context created in Validate
		maxLen := ctx.ContextData[2].(int)
		return fmt.Sprintf(defaultMaxLengthMessage, fieldValue, maxLen)
	})
}

// Validate validates the thing ensureing the field specified is populated
func (v *MaxLength) Validate(thing interface{}) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.FieldName); ok {
		if fv.Len() > v.maxLen {
			msg := specificationcatalog.Catalog().MessageStore().GetMessage(v, v.AllFieldValidators.ErrorMessageContext(thing, v.maxLen))
			return errors.NewValidationError(v.FieldName, msg)
		}
	}

	return nil
}

package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"
)

var _ interfaces.Validator = &MaxLength{}

const (
	defaultMinLengthMessage = "%v should not have a length less than %v."
	contextMinLenKey        = "MinLength"
)

// MinLength defines a validator testing the length of a field
type MinLength struct {
	*AllFieldValidators
	minLen int
}

// NewMinLengthValidator creates an initialized MaxLengthValidator
func NewMinLengthValidator(fieldName string, minLen int) interfaces.Validator {
	return &MinLength{
		AllFieldValidators: &AllFieldValidators{
			fieldName: fieldName,
		},
		minLen: minLen,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&MinLength{}, func(ctx interfaces.ValidatorContextGetter) string {
		fieldValue := ctx.GetFieldValue(ctx.GetContextData()[ContextFieldNameKey].(string))
		//nolint:errcheck // context created in Validate
		minLen := ctx.GetContextData()[contextMinLenKey].(int)
		return fmt.Sprintf(defaultMinLengthMessage, fieldValue, minLen)
	})
}

// Validate validates the thing ensuring the field specified is populated
func (v *MinLength) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.Len() < v.minLen {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				contextMaxLenKey: v.minLen,
			}))
			return NewValidationError(v.fieldName, msg)
		}
	}

	return nil
}

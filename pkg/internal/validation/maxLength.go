package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"
)

var _ interfaces.Validator = &MaxLength{}

const (
	defaultMaxLengthMessage = "%v should not have a length greater than %v."
	contextMaxLenKey        = "MaxLength"
)

// MaxLength defines a validator testing the length of a field
type MaxLength struct {
	*AllFieldValidators
	maxLen int
}

// NewMaxLengthValidator creates an initialized MaxLengthValidator
func NewMaxLengthValidator(fieldName string, maxLen int) interfaces.Validator {
	return &MaxLength{
		AllFieldValidators: &AllFieldValidators{
			fieldName: fieldName,
		},
		maxLen: maxLen,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&MaxLength{}, func(ctx interfaces.ValidatorContextGetter) string {
		fieldValue := ctx.GetFieldValue(ctx.GetContextData()[ContextFieldNameKey].(string))
		//nolint:errcheck // context created in Validate
		maxLen := ctx.GetContextData()[contextMaxLenKey].(int)
		return fmt.Sprintf(defaultMaxLengthMessage, fieldValue, maxLen)
	})
}

// Validate validates the thing ensuring the field specified is populated
func (v *MaxLength) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.Len() > v.maxLen {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				contextMaxLenKey: v.maxLen,
			}))
			return NewValidationError(v.fieldName, msg)
		}
	}

	return nil
}

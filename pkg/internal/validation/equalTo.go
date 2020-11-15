package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultEqualToMessage = "%v should be equal to %v."

// EqualTo defines a validator testing a value is Greater than another
type EqualTo struct {
	*compareValidator
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&EqualTo{}, func(ctx interfaces.ValidatorContextGetter) string {
		if compared, ok := ctx.GetContextData()[4].(bool); ok && compared {
			//nolint:errcheck // context created in Validate
			valB := ctx.GetContextData()[3]
			return fmt.Sprintf(defaultEqualToMessage, ctx.GetContextData()[0].(string), valB)
		}
		return fmt.Sprintf("Cannot compare %v to %v", ctx.GetContextData()[2], ctx.GetContextData()[3])
	})
}

// EqualToValue creates an initialized EqualTo validator comparing the value in the field to a provided value
func EqualToValue(fieldName string, equalToValue interface{}) interfaces.Validator {
	et := &EqualTo{}
	et.compareValidator = newCompareValidatorForValue(fieldName, equalToValue, []int{0}, et)

	return et
}

// EqualToFieldValue creates an initialized EqualTo validator comparing the value in the field to value in another field in the same struct
func EqualToFieldValue(fieldName, equalToFieldName string) interfaces.Validator {
	et := &EqualTo{}
	et.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, equalToFieldName, []int{0}, et)

	return et
}

// EqualToValueFromContext creates an initialized EqualTo validator comparing the value in the field to a value from the context
func EqualToValueFromContext(fieldName string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	et := &EqualTo{}
	et.compareValidator = newCompareValidatorForContext(fieldName, valueFromContext, []int{0}, et)

	return et
}

package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultGreaterThanOrEqualToMessage = "%v should be greater than or equal to %v."

// GreaterThanEqual defines a validator testing a value is GreaterThan than another
type GreaterThanEqual struct {
	*compareValidator
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&GreaterThanEqual{}, func(ctx interfaces.ValidatorContextGetter) string {
		if compared, ok := ctx.GetContextData()[4].(bool); ok && compared {
			//nolint:errcheck // context created in Validate
			valB := ctx.GetContextData()[3]
			return fmt.Sprintf(defaultGreaterThanOrEqualToMessage, ctx.GetContextData()[0].(string), valB)
		}
		return fmt.Sprintf("Cannot compare %v to %v", ctx.GetContextData()[2], ctx.GetContextData()[3])
	})
}

// GreaterThanOrEqualToValue creates an initialized GreaterThanThan validator comparing the value in the field to a provided value
func GreaterThanOrEqualToValue(fieldName string, greaterThanEqualToValue interface{}) interfaces.Validator {
	gt := &GreaterThanEqual{}
	gt.compareValidator = newCompareValidatorForValue(fieldName, greaterThanEqualToValue, []int{0, 1}, gt)

	return gt
}

// GreaterThanOrEqualToFieldValue creates an initialized GreaterThanThan validator comparing the value in the field to value in another field in the same struct
func GreaterThanOrEqualToFieldValue(fieldName, greaterThanEqualToFieldName string) interfaces.Validator {
	gt := &GreaterThanEqual{}
	gt.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, greaterThanEqualToFieldName, []int{0, 1}, gt)

	return gt
}

// GreaterThanOrEqualToValueFromContext creates an initialized GreaterThanThan validator comparing the value in the field to a value from the context
func GreaterThanOrEqualToValueFromContext(fieldName string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	gt := &GreaterThanEqual{}
	gt.compareValidator = newCompareValidatorForContext(fieldName, valueFromContext, []int{0, 1}, gt)

	return gt
}

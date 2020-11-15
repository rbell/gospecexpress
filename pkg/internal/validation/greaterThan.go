package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultGreaterThanMessage = "%v should be Greater than %v."

// GreaterThan defines a validator testing a value is Greater than another
type GreaterThan struct {
	*compareValidator
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&GreaterThan{}, func(ctx interfaces.ValidatorContextGetter) string {
		if compared, ok := ctx.GetContextData()[4].(bool); ok && compared {
			//nolint:errcheck // context created in Validate
			valB := ctx.GetContextData()[3]
			return fmt.Sprintf(defaultGreaterThanMessage, ctx.GetContextData()[0].(string), valB)
		}
		return fmt.Sprintf("Cannot compare %v to %v", ctx.GetContextData()[2], ctx.GetContextData()[3])
	})
}

// GreaterThanValue creates an initialized GreaterThan validator comparing the value in the field to a provided value
func GreaterThanValue(fieldName string, greaterThanValue interface{}) interfaces.Validator {
	gt := &GreaterThan{}
	gt.compareValidator = newCompareValidatorForValue(fieldName, greaterThanValue, []int{1}, gt)

	return gt
}

// GreaterThanFieldValue creates an initialized GreaterThan validator comparing the value in the field to value in another field in the same struct
func GreaterThanFieldValue(fieldName, greaterThanFieldName string) interfaces.Validator {
	gt := &GreaterThan{}
	gt.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, greaterThanFieldName, []int{1}, gt)

	return gt
}

// GreaterThanValueFromContext creates an initialized GreaterThan validator comparing the value in the field to a value from the context
func GreaterThanValueFromContext(fieldName string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	gt := &GreaterThan{}
	gt.compareValidator = newCompareValidatorForContext(fieldName, valueFromContext, []int{1}, gt)

	return gt
}

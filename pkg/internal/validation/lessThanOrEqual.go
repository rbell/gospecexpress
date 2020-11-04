package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultLessThanEqualToMessage = "%v should be less than or equal to %v."

// LessThanOrEqual defines a validator testing a value is less than or equal to another
type LessThanOrEqual struct {
	*compareValidator
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&LessThan{}, func(ctx interfaces.ValidatorContextGetter) string {
		if compared, ok := ctx.GetContextData()[4].(bool); ok && compared {
			//nolint:errcheck // context created in Validate
			valB := ctx.GetContextData()[3]
			return fmt.Sprintf(defaultLessThanEqualToMessage, ctx.GetContextData()[0].(string), valB)
		}
		return fmt.Sprintf("Cannot compare %v to %v", ctx.GetContextData()[2], ctx.GetContextData()[3])
	})
}

// LessThanOrEqualToValue creates an initialized MaxLengthValidator
func LessThanOrEqualToValue(fieldName string, lessThanValue interface{}) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValue(fieldName, lessThanValue, []int{-1, 0}, lt)

	return lt
}

// LessThanOrEqualToFieldValue creates an initialized LessThan validator comparing the value in the field to value in another field in the same struct
func LessThanOrEqualToFieldValue(fieldName, lessThanEqualToFieldName string) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, lessThanEqualToFieldName, []int{-1}, lt)

	return lt
}

// LessThanOrEqualToValueFromContext creates an initialized LessThan validator comparing the value in the field to a value from the context
func LessThanOrEqualToValueFromContext(fieldName string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForContext(fieldName, valueFromContext, []int{-1, 0}, lt)

	return lt
}

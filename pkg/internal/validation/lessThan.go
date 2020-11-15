package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultLessThanMessage = "%v should be less than %v."

// LessThan defines a validator testing a value is less than another
type LessThan struct {
	*compareValidator
}

func init() {
	setCompareValidatorMessage(&LessThan{}, func(ctx interfaces.ValidatorContextGetter) string {
		return fmt.Sprintf(defaultLessThanMessage, ctx.GetContextData()[ContextFieldNameKey].(string), ctx.GetContextData()[contextCompareToValueKey])
	})
}

// LessThanValue creates an initialized LessThan validator comparing the value in the field to a provided value
func LessThanValue(fieldName string, lessThanValue interface{}) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValue(fieldName, lessThanValue, []int{-1}, lt)

	return lt
}

// LessThanFieldValue creates an initialized LessThan validator comparing the value in the field to value in another field in the same struct
func LessThanFieldValue(fieldName, lessThanFieldName string) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, lessThanFieldName, []int{-1}, lt)

	return lt
}

// LessThanValueFromContext creates an initialized LessThan validator comparing the value in the field to a value from the context
func LessThanValueFromContext(fieldName string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForContext(fieldName, valueFromContext, []int{-1}, lt)

	return lt
}

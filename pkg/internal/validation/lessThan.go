package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultLessThanMessage = "%v should be less than %v."

// LessThan defines a validator testing a value is less than another
type LessThan struct {
	*compareValidator
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&LessThan{}, func(ctx interfaces.ValidatorContextGetter) string {
		if compared, ok := ctx.GetContextData()[4].(bool); ok && compared {
			//nolint:errcheck // context created in Validate
			valB := ctx.GetContextData()[3]
			return fmt.Sprintf(defaultLessThanMessage, ctx.GetContextData()[0].(string), valB)
		}
		return fmt.Sprintf("Cannot compare %v to %v", ctx.GetContextData()[2], ctx.GetContextData()[3])
	})
}

// LessThanValue creates an initialized MaxLengthValidator
func LessThanValue(fieldName string, lessThanValue interface{}) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValue(fieldName, lessThanValue, []int{-1}, lt)

	return lt
}

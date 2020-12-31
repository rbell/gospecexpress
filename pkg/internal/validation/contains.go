package validation

import (
	"fmt"
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/errors"

	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const (
	defaultContainsMessage = "%v should contain %v."
	contextContainsKey     = "Contains"
)

// Contains is a validator testing if a slice contains a value / reference
type Contains struct {
	*AllFieldValidators
	contains    interface{}
	containsVal reflect.Value
}

// NewContainsValidator returns an initialized contains validator
func NewContainsValidator(fieldName string, value interface{}) interfaces.Validator {
	return &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName: fieldName,
		},
		contains:    value,
		containsVal: reflect.ValueOf(value),
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&Contains{}, func(ctx interfaces.ValidatorContextGetter) string {
		fieldValue := ctx.GetFieldValue(ctx.GetContextData()[ContextFieldNameKey].(string))
		//nolint:errcheck // context created in Validate
		val := ctx.GetContextData()[contextMaxLenKey]
		return fmt.Sprintf(defaultContainsMessage, fieldValue, val)
	})
}

// Validate asserts that the value is found in the slice
func (v *Contains) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.Kind() == reflect.Array || fv.Kind() == reflect.Slice || fv.Kind() == reflect.String {
			for i := 0; i < fv.Len(); i++ {
				item := fv.Index(i)
				//nolint:errcheck // ignore returned error - if the values cannot be compared the item does not match the value
				eq, _ := reflectionhelpers.Eq(item, v.containsVal)
				if eq {
					// we found it - return nil
					return nil
				}
			}
		}

		// we didn't find it, return error
		msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
			contextContainsKey: v.contains,
		}))
		return errors.NewValidationError(v.fieldName, msg)
	}

	return nil
}

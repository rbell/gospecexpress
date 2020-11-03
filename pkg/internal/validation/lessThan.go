package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/compare"
)

const defaultLessThanMessage = "%v should be less than %v."

type compareFunc func(ctx *ValidatorContext) (result bool, err error)

// LessThan defines a validator testing a value is less than another
type LessThan struct {
	*AllFieldValidators
	test compareFunc
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
	lt := &LessThan{
		AllFieldValidators: &AllFieldValidators{
			FieldName: fieldName,
		},
		test: func(ctx *ValidatorContext) (result bool, err error) {
			ctx.AddContextData(ctx.GetFieldValue(fieldName))
			ctx.AddContextData(lessThanValue)
			comparer := compare.NewDefaultComparer(ctx.GetContextData()[2])
			c, err := comparer.Compare(ctx.GetContextData()[3])
			if err != nil {
				// not comparable types
				ctx.AddContextData(false)
				return false, err
			}
			// comparable types
			ctx.AddContextData(true)
			return c == -1, nil
		},
	}

	return lt
}

// Validate validates the thing ensuring the field specified is populated
func (v *LessThan) Validate(thing interface{}, messageStore interfaces.MessageStorer) error {
	ctx := v.AllFieldValidators.NewValidatorContext(thing)
	//nolint:errcheck // message store returns message if there is an error (based on context)
	if valid, _ := v.test(ctx); !valid {
		msg := messageStore.GetMessage(v, ctx)
		return NewValidationError(v.FieldName, msg)
	}
	return nil
}

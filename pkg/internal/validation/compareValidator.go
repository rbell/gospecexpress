package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/compare"
)

const (
	contextCompareToValueKey    = "CompareToValue"
	contextIsComparableTypesKey = "IsComparable"
)

type compareFunc func(ctx *ValidatorContext) (result bool, err error)

type compareValidator struct {
	*AllFieldValidators
	validatorType interfaces.Validator
	test          compareFunc
}

func setCompareValidatorMessage(validator interfaces.Validator, setter func(ctx interfaces.ValidatorContextGetter) string) {
	catalog.ValidationCatalog().MessageStore().SetMessage(validator, func(ctx interfaces.ValidatorContextGetter) string {
		if compared, ok := ctx.GetContextData()[contextIsComparableTypesKey].(bool); ok && compared {
			return setter(ctx)
		}
		return fmt.Sprintf("Cannot compare %v to %v", ctx.GetContextData()[ContextFieldValueKey], ctx.GetContextData()[contextCompareToValueKey])
	})
}

func newCompareValidatorForValue(fieldName string, value interface{}, compareValues []int, validatorType interfaces.Validator) *compareValidator {
	return newCompareValidatorForContext(fieldName, func(ctx interfaces.ValidatorContextGetter) interface{} { return value }, compareValues, validatorType)
}

func newCompareValidatorForValueAgainstOtherField(fieldName, otherFieldName string, compareValues []int, validatorType interfaces.Validator) *compareValidator {
	return newCompareValidatorForContext(fieldName, func(ctx interfaces.ValidatorContextGetter) interface{} { return ctx.GetFieldValue(otherFieldName) }, compareValues, validatorType)
}

func newCompareValidatorForContext(fieldName string, valueFromContext interfaces.ValueFromContext, compareValues []int, validatorType interfaces.Validator) *compareValidator {
	return &compareValidator{
		AllFieldValidators: &AllFieldValidators{
			FieldName: fieldName,
		},
		test: func(ctx *ValidatorContext) (result bool, err error) {
			ctx.AddContextData(ContextFieldValueKey, ctx.GetFieldValue(fieldName))
			ctx.AddContextData(contextCompareToValueKey, valueFromContext(ctx))
			comparer := compare.NewDefaultComparer(ctx.GetContextData()[ContextFieldValueKey])
			c, err := comparer.Compare(ctx.GetContextData()[contextCompareToValueKey])
			if err != nil {
				// not comparable types
				ctx.AddContextData(contextIsComparableTypesKey, false)
				return false, err
			}
			// comparable types
			ctx.AddContextData(contextIsComparableTypesKey, true)
			return intIsIn(c, compareValues), nil
		},
		validatorType: validatorType,
	}
}

func (v *compareValidator) Validate(thing interface{}, messageStore interfaces.MessageStorer) error {
	ctx := v.AllFieldValidators.NewValidatorContext(thing, nil)
	//nolint:errcheck // message store returns message if there is an error (based on context)
	if valid, _ := v.test(ctx); !valid {
		msg := messageStore.GetMessage(v.validatorType, ctx)
		return NewValidationError(v.FieldName, msg)
	}
	return nil
}

func intIsIn(intValue int, intValues []int) bool {
	for _, v := range intValues {
		if intValue == v {
			return true
		}
	}
	return false
}

package validation

import (
	"fmt"
	"reflect"
	"regexp"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	"gitlab.com/rbell/gospecexpress/pkg/errors"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"
)

const (
	defaultMatchMessage        = "%v does not match the required pattern (%v)."
	defaultMatchDescriptionKey = "MatchDescription"
)

// Match defines validator enforcing a string must match a regex
type Match struct {
	*AllFieldValidators
	regex            *regexp.Regexp
	regexDescription string
}

// NewMatch returns an initialized Match validator
func NewMatch(fieldName string, regex *regexp.Regexp, regexDescription string) interfaces.Validator {
	return &Match{
		AllFieldValidators: &AllFieldValidators{
			fieldName: fieldName,
		},
		regex:            regex,
		regexDescription: regexDescription,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&MaxLength{}, func(ctx interfaces.ValidatorContextGetter) string {
		fieldValue := ctx.GetFieldValue(ctx.GetContextData()[ContextFieldNameKey].(string))
		//nolint:errcheck // context created in Validate
		desc := ctx.GetContextData()[defaultMatchDescriptionKey].(string)
		return fmt.Sprintf(defaultMatchMessage, fieldValue, desc)
	})
}

// Validate validates the field matches the regex
func (v *Match) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok && fv.Kind() == reflect.String {
		if !v.regex.MatchString(fv.String()) {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				defaultMatchDescriptionKey: v.regexDescription,
			}))
			return errors.NewValidationError(v.fieldName, msg)
		}
	}
	return nil
}
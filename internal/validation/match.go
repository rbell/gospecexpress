// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/rbell/gospecexpress/catalog"
	"github.com/rbell/gospecexpress/errors"
	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/internal/reflectionhelpers"
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
func NewMatch(fieldName, alias string, regex *regexp.Regexp, regexDescription string) interfaces.Validator {
	return &Match{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		regex:            regex,
		regexDescription: regexDescription,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&MaxLength{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		//nolint:errcheck // ignore error
		desc := ctx.GetContextData()[defaultMatchDescriptionKey].(string)
		//nolint:errcheck // ignore error
		alias := ctx.GetContextData()[ContextFieldAliasKey].(string)
		return fmt.Sprintf(defaultMatchMessage, alias, desc)
	})
}

// Validate validates the field matches the regex
func (v *Match) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok && fv.Kind() == reflect.String {
		if !v.regex.MatchString(fv.String()) {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				defaultMatchDescriptionKey: v.regexDescription,
			}))
			return errors.NewValidationError(v.fieldName, msg, v.shouldWarn)
		}
	}
	return nil
}

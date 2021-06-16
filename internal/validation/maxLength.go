// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"github.com/rbell/gospecexpress/errors"

	"github.com/rbell/gospecexpress/catalog"

	"github.com/rbell/gospecexpress/interfaces"

	"github.com/rbell/gospecexpress/internal/reflectionhelpers"
)

var _ interfaces.Validator = &MaxLength{}

const (
	defaultMaxLengthMessage = "%v should not have a length greater than %v."
	contextMaxLenKey        = "MaxLength"
)

// MaxLength defines a validator testing the length of a field
type MaxLength struct {
	*AllFieldValidators
	maxLen int
}

// NewMaxLengthValidator creates an initialized MaxLengthValidator
func NewMaxLengthValidator(fieldName, alias string, maxLen int) interfaces.Validator {
	return &MaxLength{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		maxLen: maxLen,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&MaxLength{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		//nolint:errcheck // ignore possible error
		maxLen := ctx.GetContextData()[contextMaxLenKey].(int)
		//nolint:errcheck // ignore possible error
		alias := ctx.GetContextData()[ContextFieldAliasKey].(string)
		return fmt.Sprintf(defaultMaxLengthMessage, alias, maxLen)
	})
}

// Validate validates the thing ensuring the field specified is populated
func (v *MaxLength) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.Len() > v.maxLen {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				contextMaxLenKey: v.maxLen,
			}))
			return errors.NewValidationError(v.fieldName, msg)
		}
	}

	return nil
}

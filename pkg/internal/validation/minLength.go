// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/errors"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"
)

var _ interfaces.Validator = &MaxLength{}

const (
	defaultMinLengthMessage = "%v should not have a length less than %v."
	contextMinLenKey        = "MinLength"
)

// MinLength defines a validator testing the length of a field
type MinLength struct {
	*AllFieldValidators
	minLen int
}

// NewMinLengthValidator creates an initialized MaxLengthValidator
func NewMinLengthValidator(fieldName, alias string, minLen int) interfaces.Validator {
	return &MinLength{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		minLen: minLen,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&MinLength{}, func(ctx interfaces.ValidatorContextGetter) string {
		//nolint:errcheck // ignore possible error
		minLen := ctx.GetContextData()[contextMinLenKey].(int)
		//nolint:errcheck // ignore possible error
		alias := ctx.GetContextData()[ContextFieldAliasKey].(string)
		return fmt.Sprintf(defaultMinLengthMessage, alias, minLen)
	})
}

// Validate validates the thing ensuring the field specified is populated
func (v *MinLength) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.Len() < v.minLen {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				contextMaxLenKey: v.minLen,
			}))
			return errors.NewValidationError(v.fieldName, msg)
		}
	}

	return nil
}

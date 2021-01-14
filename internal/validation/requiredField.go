// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/errors"

	"gitlab.com/rbell/gospecexpress/catalog"

	"gitlab.com/rbell/gospecexpress/interfaces"

	"gitlab.com/rbell/gospecexpress/internal/reflectionhelpers"
)

const defaultRequiredFieldMessage = "%v is required."

// RequiredField defines a validator requiring a field value be populated.
type RequiredField struct {
	*AllFieldValidators
}

// NewRequiredFieldValidator creates an initialized RequiredFieldValidator
func NewRequiredFieldValidator(fieldName, alias string) interfaces.Validator {
	return &RequiredField{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&RequiredField{}, func(ctx interfaces.ValidatorContextGetter) string {
		return fmt.Sprintf(defaultRequiredFieldMessage, ctx.GetContextData()[ContextFieldAliasKey].(string))
	})
}

// Validate validates the thing ensureing the field specified is populated
func (v *RequiredField) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.IsZero() {
			msg := catalog.ValidationCatalog().MessageStore().GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, nil))
			return errors.NewValidationError(v.fieldName, msg)
		}
	}

	return nil
}

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
	defaultLengthEqualMessage = "%v should have a length of %v."
	contextLenEqKey           = "LengthEqualTo"
)

// LengthEquals defines a validator testing the length of a field
type LengthEquals struct {
	*AllFieldValidators
	lenEq int
}

// NewLengthEqualsValidator creates an initialized MaxLengthValidator
func NewLengthEqualsValidator(fieldName, alias string, length int) interfaces.Validator {
	return &LengthEquals{
		AllFieldValidators: &AllFieldValidators{
			fieldName:  fieldName,
			fieldAlias: alias,
		},
		lenEq: length,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&LengthEquals{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		//nolint:errcheck // context created in Validate
		maxLen := ctx.GetContextData()[contextLenEqKey].(int)
		//nolint:errcheck // context created in Validate
		alias := ctx.GetContextData()[ContextFieldAliasKey].(string)
		return fmt.Sprintf(defaultLengthEqualMessage, alias, maxLen)
	})
}

// Validate validates the thing ensuring the field specified is populated
func (v *LengthEquals) Validate(thing interface{}, contextData map[string]interface{}, messageStore interfaces.MessageStorer) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.Len() != v.lenEq {
			msg := messageStore.GetMessage(v, v.AllFieldValidators.NewValidatorContext(thing, map[string]interface{}{
				contextLenEqKey: v.lenEq,
			}))
			return errors.NewValidationError(v.fieldName, msg)
		}
	}

	return nil
}

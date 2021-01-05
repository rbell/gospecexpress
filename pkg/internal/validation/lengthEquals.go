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
	defaultLengthEqualMessage = "%v should have a length of %v."
	contextLenEqKey           = "LengthEqualTo"
)

// LengthEquals defines a validator testing the length of a field
type LengthEquals struct {
	*AllFieldValidators
	lenEq int
}

// NewLengthEqualsValidator creates an initialized MaxLengthValidator
func NewLengthEqualsValidator(fieldName string, length int) interfaces.Validator {
	return &LengthEquals{
		AllFieldValidators: &AllFieldValidators{
			fieldName: fieldName,
		},
		lenEq: length,
	}
}

func init() {
	catalog.ValidationCatalog().MessageStore().SetMessage(&LengthEquals{}, func(ctx interfaces.ValidatorContextGetter) string {
		fieldValue := ctx.GetFieldValue(ctx.GetContextData()[ContextFieldNameKey].(string))
		//nolint:errcheck // context created in Validate
		maxLen := ctx.GetContextData()[contextLenEqKey].(int)
		return fmt.Sprintf(defaultLengthEqualMessage, fieldValue, maxLen)
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

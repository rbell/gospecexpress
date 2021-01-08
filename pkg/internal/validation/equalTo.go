// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultEqualToMessage = "%v should be equal to %v."

// EqualTo defines a validator testing a value is Greater than another
type EqualTo struct {
	*compareValidator
}

func init() {
	setCompareValidatorMessage(&EqualTo{}, func(ctx interfaces.ValidatorContextGetter) string {
		return fmt.Sprintf(defaultEqualToMessage, ctx.GetContextData()[ContextFieldAliasKey].(string), ctx.GetContextData()[contextCompareToValueKey])
	})
}

// EqualToValue creates an initialized EqualTo validator comparing the value in the field to a provided value
func EqualToValue(fieldName, alias string, equalToValue interface{}) interfaces.Validator {
	et := &EqualTo{}
	et.compareValidator = newCompareValidatorForValue(fieldName, alias, equalToValue, []int{0}, et)

	return et
}

// EqualToFieldValue creates an initialized EqualTo validator comparing the value in the field to value in another field in the same struct
func EqualToFieldValue(fieldName, alias, equalToFieldName string) interfaces.Validator {
	et := &EqualTo{}
	et.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, alias, equalToFieldName, []int{0}, et)

	return et
}

// EqualToValueFromContext creates an initialized EqualTo validator comparing the value in the field to a value from the context
func EqualToValueFromContext(fieldName, alias string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	et := &EqualTo{}
	et.compareValidator = newCompareValidatorForContext(fieldName, alias, et, &valueCompare{
		getValue:            valueFromContext,
		compareToContextKey: contextCompareToValueKey,
		comparisonValues:    []int{0},
	})

	return et
}

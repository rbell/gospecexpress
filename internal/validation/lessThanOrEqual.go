// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"github.com/rbell/gospecexpress/interfaces"
)

const defaultLessThanEqualToMessage = "%v should be less than or equal to %v."

// LessThanOrEqual defines a validator testing a value is less than or equal to another
type LessThanOrEqual struct {
	*compareValidator
}

func init() {
	setCompareValidatorMessage(&LessThanOrEqual{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		return fmt.Sprintf(defaultLessThanEqualToMessage, ctx.GetContextData()[ContextFieldAliasKey].(string), ctx.GetContextData()[contextCompareToValueKey])
	})
}

// LessThanOrEqualToValue creates an initialized MaxLengthValidator
func LessThanOrEqualToValue(fieldName, alias string, lessThanValue interface{}) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValue(fieldName, alias, lessThanValue, []int{-1, 0}, lt)

	return lt
}

// LessThanOrEqualToFieldValue creates an initialized LessThan validator comparing the value in the field to value in another field in the same struct
func LessThanOrEqualToFieldValue(fieldName, alias, lessThanEqualToFieldName string) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, alias, lessThanEqualToFieldName, []int{-1}, lt)

	return lt
}

// LessThanOrEqualToValueFromContext creates an initialized LessThan validator comparing the value in the field to a value from the context
func LessThanOrEqualToValueFromContext(fieldName, alias string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForContext(fieldName, alias, lt, &valueCompare{
		getValue:            valueFromContext,
		compareToContextKey: contextCompareToValueKey,
		comparisonValues:    []int{-1, 0},
	})

	return lt
}

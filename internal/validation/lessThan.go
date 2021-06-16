// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"github.com/rbell/gospecexpress/interfaces"
)

const defaultLessThanMessage = "%v should be less than %v."

// LessThan defines a validator testing a value is less than another
type LessThan struct {
	*compareValidator
}

func init() {
	setCompareValidatorMessage(&LessThan{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		return fmt.Sprintf(defaultLessThanMessage, ctx.GetContextData()[ContextFieldAliasKey].(string), ctx.GetContextData()[contextCompareToValueKey])
	})
}

// LessThanValue creates an initialized LessThan validator comparing the value in the field to a provided value
func LessThanValue(fieldName, alias string, lessThanValue interface{}) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValue(fieldName, alias, lessThanValue, []int{-1}, lt)

	return lt
}

// LessThanFieldValue creates an initialized LessThan validator comparing the value in the field to value in another field in the same struct
func LessThanFieldValue(fieldName, alias, lessThanFieldName string) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, alias, lessThanFieldName, []int{-1}, lt)

	return lt
}

// LessThanValueFromContext creates an initialized LessThan validator comparing the value in the field to a value from the context
func LessThanValueFromContext(fieldName, alias string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	lt := &LessThan{}
	lt.compareValidator = newCompareValidatorForContext(fieldName, alias, lt, &valueCompare{
		getValue:            valueFromContext,
		compareToContextKey: contextCompareToValueKey,
		comparisonValues:    []int{-1},
	})

	return lt
}

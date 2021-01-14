// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"github.com/rbell/gospecexpress/interfaces"
)

const defaultGreaterThanOrEqualToMessage = "%v should be greater than or equal to %v."

// GreaterThanEqual defines a validator testing a value is GreaterThan than another
type GreaterThanEqual struct {
	*compareValidator
}

func init() {
	setCompareValidatorMessage(&GreaterThanEqual{}, func(ctx interfaces.ValidatorContextGetter) string {
		return fmt.Sprintf(defaultGreaterThanOrEqualToMessage, ctx.GetContextData()[ContextFieldAliasKey].(string), ctx.GetContextData()[contextCompareToValueKey])
	})
}

// GreaterThanOrEqualToValue creates an initialized GreaterThanThan validator comparing the value in the field to a provided value
func GreaterThanOrEqualToValue(fieldName, alias string, greaterThanEqualToValue interface{}) interfaces.Validator {
	gt := &GreaterThanEqual{}
	gt.compareValidator = newCompareValidatorForValue(fieldName, alias, greaterThanEqualToValue, []int{0, 1}, gt)

	return gt
}

// GreaterThanOrEqualToFieldValue creates an initialized GreaterThanThan validator comparing the value in the field to value in another field in the same struct
func GreaterThanOrEqualToFieldValue(fieldName, alias, greaterThanEqualToFieldName string) interfaces.Validator {
	gt := &GreaterThanEqual{}
	gt.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, alias, greaterThanEqualToFieldName, []int{0, 1}, gt)

	return gt
}

// GreaterThanOrEqualToValueFromContext creates an initialized GreaterThanThan validator comparing the value in the field to a value from the context
func GreaterThanOrEqualToValueFromContext(fieldName, alias string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	gt := &GreaterThanEqual{}
	gt.compareValidator = newCompareValidatorForContext(fieldName, alias, gt, &valueCompare{
		getValue:            valueFromContext,
		compareToContextKey: contextCompareToValueKey,
		comparisonValues:    []int{0, 1},
	})

	return gt
}

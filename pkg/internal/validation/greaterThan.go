// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultGreaterThanMessage = "%v should be Greater than %v."

// GreaterThan defines a validator testing a value is Greater than another
type GreaterThan struct {
	*compareValidator
}

func init() {
	setCompareValidatorMessage(&GreaterThan{}, func(ctx interfaces.ValidatorContextGetter) string {
		return fmt.Sprintf(defaultGreaterThanMessage, ctx.GetContextData()[ContextFieldNameKey].(string), ctx.GetContextData()[contextCompareToValueKey])
	})
}

// GreaterThanValue creates an initialized GreaterThan validator comparing the value in the field to a provided value
func GreaterThanValue(fieldName string, greaterThanValue interface{}) interfaces.Validator {
	gt := &GreaterThan{}
	gt.compareValidator = newCompareValidatorForValue(fieldName, greaterThanValue, []int{1}, gt)

	return gt
}

// GreaterThanFieldValue creates an initialized GreaterThan validator comparing the value in the field to value in another field in the same struct
func GreaterThanFieldValue(fieldName, greaterThanFieldName string) interfaces.Validator {
	gt := &GreaterThan{}
	gt.compareValidator = newCompareValidatorForValueAgainstOtherField(fieldName, greaterThanFieldName, []int{1}, gt)

	return gt
}

// GreaterThanValueFromContext creates an initialized GreaterThan validator comparing the value in the field to a value from the context
func GreaterThanValueFromContext(fieldName string, valueFromContext interfaces.ValueFromContext) interfaces.Validator {
	gt := &GreaterThan{}
	gt.compareValidator = newCompareValidatorForContext(fieldName, valueFromContext, []int{1}, gt)

	return gt
}

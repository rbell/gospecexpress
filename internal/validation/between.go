// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"

	"github.com/rbell/gospecexpress/interfaces"
)

const (
	defaultBetweenMessage = "%v should be between %v and %v."
	lowerContextKey       = "Lower"
	upperContextKey       = "upper"
)

// Between defines a validator testing that a value is between two other values
type Between struct {
	*compareValidator
}

func init() {
	setCompareValidatorMessage(&Between{}, func(ctx interfaces.FieldValidatorContextGetter) string {
		return fmt.Sprintf(defaultBetweenMessage, ctx.GetContextData()[ContextFieldAliasKey], ctx.GetContextData()[lowerContextKey], ctx.GetContextData()[upperContextKey])
	})
}

// BetweenValues creates an initialized Between validator ensuring the value in the field is between the lower and upper parameters
func BetweenValues(fieldName, alias string, lower, upper interface{}) interfaces.Validator {
	return BetweenValuesFromContext(fieldName, alias,
		func(ctx interfaces.FieldValidatorContextGetter) interface{} { return lower },
		func(ctx interfaces.FieldValidatorContextGetter) interface{} { return upper },
	)
}

// BetweenOtherFieldValues creates an initialized Between validator ensuring the value in the field is between values stored in two other fields
func BetweenOtherFieldValues(fieldName, alias, lowerFieldName, upperFieldName string) interfaces.Validator {
	return BetweenValuesFromContext(fieldName, alias,
		func(ctx interfaces.FieldValidatorContextGetter) interface{} { return ctx.GetFieldValue(lowerFieldName) },
		func(ctx interfaces.FieldValidatorContextGetter) interface{} { return ctx.GetFieldValue(upperFieldName) },
	)
}

// BetweenValuesFromContext creates an initialized Between validator ensuring the value in the field is between values retrieved from validation context
func BetweenValuesFromContext(fieldName, alias string, lowerGetter, upperGetter interfaces.ValueFromContext) interfaces.Validator {
	between := &Between{}
	between.compareValidator = newCompareValidatorForContext(fieldName, alias, between,
		&valueCompare{
			getValue:            lowerGetter,
			comparisonValues:    []int{1},
			compareToContextKey: lowerContextKey,
		},
		&valueCompare{
			getValue:            upperGetter,
			comparisonValues:    []int{-1},
			compareToContextKey: upperContextKey,
		},
	)
	return between
}

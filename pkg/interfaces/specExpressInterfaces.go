// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interfaces

import (
	"reflect"
)

// ValueFromContext defines functor returning a value from a ValidatorContext
type ValueFromContext func(ctx ValidatorContextGetter) interface{}

// SpecificationValidator defines interface to Validate something
type SpecificationValidator interface {
	Validate(subject interface{}, contextData map[string]interface{}) error
	ForType(forType interface{}) QualifierBuilder
	GetForType() reflect.Type
}

// QualifierBuilder defines interface for starting to qualify an element
type QualifierBuilder interface {
	Required(fieldName string, options ...ValidatorOption) ValidatorBuilder
}

// ValidatorBuilder defines interface methods to build a specification
type ValidatorBuilder interface {
	// Qualifier Validation
	Required(fieldName string, options ...ValidatorOption) ValidatorBuilder
	// TODO: Optional

	// String Specific Validators
	// TODO: Matches

	// Date Validators
	// TODO: Before
	// TODO: After
	// TODO: BeforeOtherField
	// TODO: AfterOtherField
	// TODO: BeforeValueFromContext
	// TODO: AfterValueFromContext

	// Compare Validators
	// TODO: Between
	// TODO: BetweenOtherFields
	// TODO: BetweenValuesFromContext
	LessThan(value interface{}, options ...ValidatorOption) ValidatorBuilder
	LessThanOtherField(otherField string, options ...ValidatorOption) ValidatorBuilder
	LessThanValueFromContext(valueFromContext ValueFromContext, options ...ValidatorOption) ValidatorBuilder
	LessThanOrEqualTo(value interface{}, options ...ValidatorOption) ValidatorBuilder
	LessThanOrEqualToOtherField(otherField string, options ...ValidatorOption) ValidatorBuilder
	LessThanOrEqualToValueFromContext(valueFromContext ValueFromContext, options ...ValidatorOption) ValidatorBuilder
	GreaterThan(value interface{}, options ...ValidatorOption) ValidatorBuilder
	GreaterThanOtherField(otherField string, options ...ValidatorOption) ValidatorBuilder
	GreaterThanValueFromContext(valueFromContext ValueFromContext, options ...ValidatorOption) ValidatorBuilder
	GreaterThanOrEqualTo(value interface{}, options ...ValidatorOption) ValidatorBuilder
	GreaterThanOrEqualToOtherField(otherField string, options ...ValidatorOption) ValidatorBuilder
	GreaterThanOrEqualToValueFromContext(valueFromContext ValueFromContext, options ...ValidatorOption) ValidatorBuilder
	EqualTo(value interface{}, options ...ValidatorOption) ValidatorBuilder
	EqualToOtherField(otherField string, options ...ValidatorOption) ValidatorBuilder
	EqualToValueFromContext(valueFromContext ValueFromContext, options ...ValidatorOption) ValidatorBuilder

	// Slice Validators (strings are considered slices)
	MaxLength(len int) ValidatorBuilder
	MinLength(len int) ValidatorBuilder
	Contains(thing interface{}) ValidatorBuilder
	ContainsValueFromContext(fromContext ValueFromContext) ValidatorBuilder
	// TODO: CountEqual
	// TODO: RangeValidate
	// TODO: RangeExpect

	// Reference Validators
	ValidateReference() ValidatorBuilder

	// Custom Rule which if returned error is not nil, error's message will be included in the validation error
	Expect(validator func(validationCtx ValidatorContextGetter) error) ValidatorBuilder
}

// Validator defines interface for something that can validate.  Similar to a boolean predicate, a validator returns
type Validator interface {
	Validate(thing interface{}, contextData map[string]interface{}, messageStore MessageStorer) error
}

// ValidatorOption is a function signature defining an option on a Validator
type ValidatorOption func(v Validator)

// MessageFormatter defines a function that returns a message given a ValidatorContextGetter
type MessageFormatter func(ctx ValidatorContextGetter) string

// MessageOverrider defines interface for something that has ability to override a validation message
type MessageOverrider interface {
	GetOverrideErrorMessage(ctx ValidatorContextGetter) string
	SetOverrideErrorMessage(msgFormatter MessageFormatter)
}

// MessageStorer defines interface for getting a message for a validation rule
type MessageStorer interface {
	GetMessage(validator Validator, ctx ValidatorContextGetter) string
	SetMessage(validator Validator, getterFunc ErrorMessageGetterFunc)
}

// ValidatorContextGetter gets the context for the validation
type ValidatorContextGetter interface {
	GetFieldValue(fieldName string) interface{}
	// ContextData will include
	GetContextData() map[string]interface{}
}

// Cataloger defines interface for a validation catalog
type Cataloger interface {
	Register(s SpecificationValidator)
	Validate(something interface{}) error
	ValidateWithContext(something interface{}, contextData map[string]interface{}) error
	MessageStore() MessageStorer
}

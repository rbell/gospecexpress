// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interfaces

import (
	"reflect"
	"regexp"
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
	Optional(fieldName string) ValidatorBuilder
}

// ValidatorBuilder defines interface methods to build a specification
type ValidatorBuilder interface {
	// Qualifier Validation
	Required(fieldName string, options ...ValidatorOption) ValidatorBuilder
	Optional(fieldName string) ValidatorBuilder

	// String Specific Validators
	Matches(regex *regexp.Regexp, regexDescripton string, options ...ValidatorOption) ValidatorBuilder

	// Compare Validators
	Between(lower, upper interface{}, options ...ValidatorOption) ValidatorBuilder
	BetweenOtherFields(lowerField, upperField string, options ...ValidatorOption) ValidatorBuilder
	BetweenValuesFromContext(lowerGetter, upperGetter ValueFromContext, options ...ValidatorOption) ValidatorBuilder
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
	LengthEquals(length int, options ...ValidatorOption) ValidatorBuilder
	MaxLength(len int, options ...ValidatorOption) ValidatorBuilder
	MinLength(len int, options ...ValidatorOption) ValidatorBuilder
	Contains(thing interface{}, options ...ValidatorOption) ValidatorBuilder
	ContainsValueFromContext(fromContext ValueFromContext, options ...ValidatorOption) ValidatorBuilder
	RangeValidate(options ...ValidatorOption) ValidatorBuilder
	RangeExpect(validator func(validationCtx ValidatorContextGetter) error, options ...ValidatorOption) ValidatorBuilder

	// Reference Validators
	ValidateReference(options ...ValidatorOption) ValidatorBuilder

	// Custom Rule which if returned error is not nil, error's message will be included in the validation error
	Expect(validator func(validationCtx ValidatorContextGetter) error, options ...ValidatorOption) ValidatorBuilder
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

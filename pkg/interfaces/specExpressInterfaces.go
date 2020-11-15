package interfaces

import (
	"reflect"
)

// ValueFromContext defines functor returning a value from a ValidatorContext
type ValueFromContext func(ctx ValidatorContextGetter) interface{}

// SpecificationValidator defines interface to Validate something
type SpecificationValidator interface {
	Validate(interface{}) error
	ForType(forType interface{}) QualifierBuilder
	GetForType() reflect.Type
}

// QualifierBuilder defines interface for starting to qualify an element
type QualifierBuilder interface {
	RequiredField(fieldName string) ValidatorBuilder
}

// ValidatorBuilder defines interface methods to build a specification
type ValidatorBuilder interface {
	RequiredField(fieldName string) ValidatorBuilder

	// String Validators
	MaxLength(len int) ValidatorBuilder

	// Compare Validators
	LessThan(value interface{}) ValidatorBuilder
	LessThanOtherField(otherField string) ValidatorBuilder
	LessThanValueFromContext(valueFromContext ValueFromContext) ValidatorBuilder
	LessThanOrEqualTo(value interface{}) ValidatorBuilder
	LessThanOrEqualToOtherField(otherField string) ValidatorBuilder
	LessThanOrEqualToValueFromContext(valueFromContext ValueFromContext) ValidatorBuilder
	GreaterThan(value interface{}) ValidatorBuilder
	GreaterThanOtherField(otherField string) ValidatorBuilder
	GreaterThanValueFromContext(valueFromContext ValueFromContext) ValidatorBuilder
	GreaterThanOrEqualTo(value interface{}) ValidatorBuilder
	GreaterThanOrEqualToOtherField(otherField string) ValidatorBuilder
	GreaterThanOrEqualToValueFromContext(valueFromContext ValueFromContext) ValidatorBuilder
	EqualTo(value interface{}) ValidatorBuilder
	EqualToOtherField(otherField string) ValidatorBuilder
	EqualToValueFromContext(valueFromContext ValueFromContext) ValidatorBuilder
}

// Validator defines interface for something that can validate.  Similar to a boolean predicate, a validator returns
type Validator interface {
	Validate(thing interface{}, messageStore MessageStorer) error
}

// MessageStorer defines interface for getting a message for a validation rule
type MessageStorer interface {
	GetMessage(validator Validator, ctx ValidatorContextGetter) string
	SetMessage(validator Validator, getterFunc ErrorMessageGetterFunc)
}

// ValidatorContextGetter gets the context for the validation
type ValidatorContextGetter interface {
	GetFieldValue(fieldName string) interface{}
	GetContextData() map[string]interface{}
}

// Cataloger defines interface for a validation catalog
type Cataloger interface {
	Register(s SpecificationValidator)
	Validate(something interface{}) error
	MessageStore() MessageStorer
}

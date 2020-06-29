package interfaces

import "reflect"

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
}

// Validator defines interface for something that can validate.  Similar to a boolean predicate, a validator returns
type Validator interface {
	Validate(thing interface{}) error
}

// MessageStorer defines interface for getting a message for a validation rule
type MessageStorer interface {
	GetMessage(validator Validator) string
	StoreMessage(validator Validator, msg string)
}

// Cataloger defines interface for a validation catalog
type Cataloger interface {
	Register(s SpecificationValidator)
	Validate(something interface{}) error
	MessageStore() MessageStorer
}

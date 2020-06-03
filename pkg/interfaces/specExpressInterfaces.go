package interfaces

import "reflect"

// SpecificationValidator defines interface to Validate something
type SpecificationValidator interface {
	Validate(interface{}) error
	ForType(forType interface{}) ValidatorBuilder
	GetForType() reflect.Type
}

// ValidatorBuilder defines interface methods to build a specification
type ValidatorBuilder interface {
	RequiredField(fieldName string) ValidatorBuilder
}

// Validator defines interface for something that can validate
type Validator interface {
	Validate(thing interface{}) error
}

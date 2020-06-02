package interfaces

import "reflect"

// SpecificationValidator defines interface to Validate something
type SpecificationValidator interface {
	Validate(interface{}) error
	ForType(forType interface{}) ValidatorBuilder
	GetForType() reflect.Type
}

// SpecificationBuilder defines interface methods to build a specification
type ValidatorBuilder interface {
	RequiredField(fieldName string) ValidatorBuilder
}

type Validator interface {
	Validate(thing interface{}) error
}

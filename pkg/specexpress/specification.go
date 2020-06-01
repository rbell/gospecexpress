package specexpress

import "reflect"

// SpecificationValidator defines interface to Validate something
type SpecificationValidator interface {
	Validate(interface{}) bool
}

// SpecificationBuilder defines interface methods to build a specification
type SpecificationBuilder interface {
	SpecificationValidator
	ForType(forType interface{}) SpecificationBuilder
	GetForType() reflect.Type
}

// Specification defines a base for specification
type Specification struct {
	forType reflect.Type
}

// ForType sets the type that the specification is to be applied to
func (s *Specification) ForType(forType interface{}) SpecificationBuilder {
	s.forType = reflect.TypeOf(forType)
	return s
}

// GetForType returns the type that the specification is to be applied to
func (s *Specification) GetForType() reflect.Type {
	return s.forType
}

// Validate validates an instance of the type
func (s *Specification) Validate(interface{}) bool {
	return true
}

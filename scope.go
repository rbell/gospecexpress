package gospecexpress

import "github.com/rbell/gospecexpress/interfaces"

var _ interfaces.SpecificationScoper = &specificationScope{}

type specificationScope struct {
	scopeName                   string
	extendsDefaultSpecification bool
}

// NewSpecificationScope returns a Specification scope
func NewSpecificationScope(name string, extendsDefaultSpecification bool) interfaces.SpecificationScoper {
	return &specificationScope{scopeName: name, extendsDefaultSpecification: extendsDefaultSpecification}
}

func (s *specificationScope) GetScopeName() string {
	return s.scopeName
}

func (s *specificationScope) ExtendsDefaultSpecification() bool {
	return s.extendsDefaultSpecification
}

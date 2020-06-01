package specexpress

import (
	"reflect"
	"sync"
)

const (
	defaultContext = "default"
)

// Cataloger defines interface for a validation catalog
type Cataloger interface {
	Register(s SpecificationBuilder)
	Validate(something interface{}) bool
}

var instance Cataloger
var instanceOnce = &sync.Once{}

type catalog struct {
	validators map[reflect.Type]map[string]SpecificationBuilder
}

// Catalog gets the singleton instance of the Cataloger
func Catalog() Cataloger {
	instanceOnce.Do(func() {
		instance = &catalog{
			validators: make(map[reflect.Type]map[string]SpecificationBuilder),
		}
	})

	return instance
}

// Register registers a specification in the catalog
func (c *catalog) Register(s SpecificationBuilder) {
	t := s.GetForType()
	if c.validators[t] == nil {
		c.validators[t] = make(map[string]SpecificationBuilder)
	}
	c.validators[t][defaultContext] = s
}

// Validate validates something against the catalog of specifications
func (c *catalog) Validate(something interface{}) bool {
	t := reflect.TypeOf(something)
	if vs,ok := c.validators[t]; ok {
		if v,ok := vs[defaultContext];ok {
			return v.Validate(something)
		}
	}
	return false
}

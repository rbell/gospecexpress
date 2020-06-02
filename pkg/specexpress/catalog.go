package specexpress

import (
	"gitlab.com/govalidate/pkg/interfaces"
	"reflect"
	"sync"
)

const (
	defaultContext = "default"
)

// Cataloger defines interface for a validation catalog
type Cataloger interface {
	Register(s interfaces.SpecificationValidator)
	Validate(something interface{}) error
}

var instance Cataloger
var instanceOnce = &sync.Once{}

type catalog struct {
	validators map[reflect.Type]map[string]interfaces.SpecificationValidator
}

// Catalog gets the singleton instance of the Cataloger
func Catalog() Cataloger {
	instanceOnce.Do(func() {
		instance = &catalog{
			validators: make(map[reflect.Type]map[string]interfaces.SpecificationValidator),
		}
	})

	return instance
}

// Register registers a specification in the catalog
func (c *catalog) Register(s interfaces.SpecificationValidator) {
	t := s.GetForType()
	if c.validators[t] == nil {
		c.validators[t] = make(map[string]interfaces.SpecificationValidator)
	}
	c.validators[t][defaultContext] = s
}

// Validate validates something against the catalog of specifications
func (c *catalog) Validate(something interface{}) error {
	t := reflect.TypeOf(something)
	if vs,ok := c.validators[t]; ok {
		if v,ok := vs[defaultContext];ok {
			return v.Validate(something)
		}
	}

	// Catalog does not contain specification for something or it is valid.
	return nil
}

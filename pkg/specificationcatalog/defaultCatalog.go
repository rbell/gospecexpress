package specificationcatalog

import (
	"reflect"
	"sync"

	"gitlab.com/rbell/gospecexpress/pkg/errormessagestore"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const (
	defaultContext = "default"
)

var instance interfaces.Cataloger
var instanceOnce = &sync.Once{}

// Catalog gets the singleton instance of the Cataloger
func Catalog() interfaces.Cataloger {
	instanceOnce.Do(func() {
		instance = &catalog{
			validators:   make(map[reflect.Type]map[string]interfaces.SpecificationValidator),
			messageStore: errormessagestore.NewDefaultMessageStore(),
		}
	})

	return instance
}

type catalog struct {
	validators   map[reflect.Type]map[string]interfaces.SpecificationValidator
	messageStore interfaces.MessageStorer
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
	if vs, ok := c.validators[t]; ok {
		if v, ok := vs[defaultContext]; ok {
			return v.Validate(something)
		}
	}

	// Catalog does not contain specification for something or it is valid.
	return nil
}

// Message Store returns the currently configured MessageStore
func (c *catalog) MessageStore() interfaces.MessageStorer {
	return c.messageStore
}

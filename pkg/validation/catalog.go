package validation

import (
	"reflect"
	"sync"
)

const (
	defaultContext = "default"
)

type Cataloger interface {
	RegisterForType(forType interface{}, v Validator)
	Validate(something interface{}) bool
}

var instance Cataloger
var instanceOnce = &sync.Once{}

type catalog struct {
	validators map[reflect.Type]map[string]Validator
}

func Catalog() Cataloger {
	instanceOnce.Do(func() {
		instance = &catalog{
			validators: make(map[reflect.Type]map[string]Validator),
		}
	})

	return instance
}

func (c *catalog) RegisterForType(forType interface{}, v Validator) {
	t := reflect.TypeOf(forType)
	if c.validators[t] == nil {
		c.validators[t] = make(map[string]Validator)
	}
	c.validators[t][defaultContext] = v
}

func (c *catalog) Validate(something interface{}) bool {
	t := reflect.TypeOf(something)
	if vs,ok := c.validators[t]; ok {
		if v,ok := vs[defaultContext];ok {
			return v.IsValid()
		}
	}
	return false
}

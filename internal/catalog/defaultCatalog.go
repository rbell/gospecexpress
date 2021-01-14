// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package catalog

import (
	"reflect"

	"github.com/rbell/gospecexpress/interfaces"
)

const (
	defaultContext = "default"
)

// DefaultCatalog is the default validation catalog supported for specexpress
type DefaultCatalog struct {
	validators   map[reflect.Type]map[string]interfaces.SpecificationValidator
	messageStore interfaces.MessageStorer
}

// NewDefaultCatalog gets an initialized default catalog
func NewDefaultCatalog() *DefaultCatalog {
	return &DefaultCatalog{
		validators:   make(map[reflect.Type]map[string]interfaces.SpecificationValidator),
		messageStore: NewDefaultMessageStore(),
	}
}

// Register registers a specification in the DefaultCatalog
func (c *DefaultCatalog) Register(s interfaces.SpecificationValidator) {
	t := s.GetForType()
	if c.validators[t] == nil {
		c.validators[t] = make(map[string]interfaces.SpecificationValidator)
	}
	c.validators[t][defaultContext] = s
}

// Validate validates something against the DefaultCatalog of specifications
func (c *DefaultCatalog) Validate(something interface{}) error {
	return c.ValidateWithContext(something, nil)
}

// ValidateWithContext validates something against the DefaultCatalog, with additional context to be used in the validation
// The additional context is a map which can be referenced by the registered validators associated with the subject in the catalog
func (c *DefaultCatalog) ValidateWithContext(something interface{}, contextData map[string]interface{}) error {
	t := reflect.TypeOf(something)
	if vs, ok := c.validators[t]; ok {
		if v, ok := vs[defaultContext]; ok {
			return v.Validate(something, contextData)
		}
	}

	// Catalog does not contain specification for something or it is valid.
	return nil
}

// MessageStore returns the currently configured MessageStore
func (c *DefaultCatalog) MessageStore() interfaces.MessageStorer {
	return c.messageStore
}

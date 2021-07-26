// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package catalog

import (
	"fmt"
	"reflect"

	"github.com/rbell/gospecexpress/errors"
	"github.com/rbell/gospecexpress/internal/errorhelpers"

	"github.com/rbell/gospecexpress/interfaces"
)

const (
	defaultScope = "default"
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
	scopeKey := defaultScope
	scope := s.GetScope()
	if scope != nil && scope.GetScopeName() != "" {
		scopeKey = scope.GetScopeName()
	}
	if c.validators[t] == nil {
		c.validators[t] = make(map[string]interfaces.SpecificationValidator)
	}
	c.validators[t][scopeKey] = s
}

// Validate validates something against the DefaultCatalog of specifications
func (c *DefaultCatalog) Validate(something interface{}) error {
	return c.ValidateWithContext(something, nil)
}

// ValidateWithContext validates something against the DefaultCatalog, with additional context to be used in the validation
// The additional context is a map which can be referenced by the registered validators associated with the subject in the catalog
func (c *DefaultCatalog) ValidateWithContext(something interface{}, contextData map[string]interface{}) error {
	t := reflect.TypeOf(something)
	scopeKey := defaultScope
	if specificScope, ok := contextData[interfaces.ScopeContextKey]; ok {
		//nolint:errcheck // ignore checking error since we know its a string
		scopeKey = specificScope.(string)
	}

	var valError *errors.ValidatorError = nil
	if vs, ok := c.validators[t]; ok {

		var defaultSpec, scopedSpec interfaces.SpecificationValidator
		extendsDefaultSpec := false

		// locate scopedSpec if validating for a scope
		if scopeKey != defaultScope {
			if v, ok := vs[scopeKey]; ok {
				scopedSpec = v
				if scope := scopedSpec.GetScope(); scope != nil {
					extendsDefaultSpec = scope.ExtendsDefaultSpecification()
				}
			} else {
				// A scoped specification was requested but not found.  Return error indicating such.
				//nolint:golint // Ignore suggestion of "error strings should not be capitalized or end with punctuation or a newline" - needs to be user readable
				return fmt.Errorf("There is no specification for %v registered in the catalog for the %v scopeKey.", t.String(), scopeKey)
			}
		}

		// locate default spec if needed
		if extendsDefaultSpec || scopeKey == defaultScope {
			if v, ok := vs[defaultScope]; ok {
				defaultSpec = v
			}
		}

		// Now validate default spec if we have it
		if defaultSpec != nil {
			e := defaultSpec.Validate(something, contextData)
			if e != nil {
				if ve, ok := errors.IsValidatorError(e); ok {
					valError = ve
				} else {
					return e
				}
			}
		}

		// Add in the results of the scopedSpec
		if scopedSpec != nil {
			e := scopedSpec.Validate(something, contextData)
			if e != nil {
				if ve, ok := errors.IsValidatorError(e); ok {
					valError = errorhelpers.JoinErrors(valError, ve)
				} else {
					return e
				}
			}
		}
	}

	// Catalog does not contain specification for something or it is valid.
	if valError != nil {
		return valError
	}

	return nil
}

// MessageStore returns the currently configured MessageStore
func (c *DefaultCatalog) MessageStore() interfaces.MessageStorer {
	return c.messageStore
}

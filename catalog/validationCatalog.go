// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package catalog

import (
	"sync"

	"github.com/rbell/gospecexpress/internal/catalog"

	"github.com/rbell/gospecexpress/interfaces"
)

var instance interfaces.Cataloger
var instanceOnce = &sync.Once{}

// ValidationCatalog gets the singleton instance of the Cataloger
func ValidationCatalog() interfaces.Cataloger {
	instanceOnce.Do(func() {
		instance = catalog.NewDefaultCatalog()
	})

	return instance
}

// Validate validates an instance against the default validation catalog
func Validate(something interface{}, options ...interfaces.ValidateOption) error {
	if len(options) > 0 {
		contextData := make(map[string]interface{})
		for _, opt := range options {
			opt(something, contextData)
		}
		return ValidationCatalog().ValidateWithContext(something, contextData)
	}
	return ValidationCatalog().Validate(something)
}

// Register a specification into the default validation catalog
func Register(spec interfaces.SpecificationValidator) {
	ValidationCatalog().Register(spec)
}

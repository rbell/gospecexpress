// Copyright ©2021 by Randy R Bell. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package catalog

import (
	"reflect"
	"testing"

	"github.com/rbell/gospecexpress/interfaces"

	"github.com/rbell/gospecexpress/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCatalog_RegisterForType_ShouldRegisterDefalutSpecForType(t *testing.T) {
	// setup
	c := &DefaultCatalog{validators: make(map[reflect.Type]map[string]interfaces.SpecificationValidator)}
	mSpec := &mocks.SpecificationValidator{}
	type fakeStruct struct{}
	fake := &fakeStruct{}
	fakeType := reflect.TypeOf(fake)

	mSpec.On("GetForType").Return(fakeType)
	mSpec.On("GetScope").Return(nil)

	// test
	c.Register(mSpec)

	// assert
	mSpec.AssertExpectations(t)
	assert.Contains(t, c.validators, fakeType)
	assert.Contains(t, c.validators[fakeType], defaultScope)
	assert.Equal(t, mSpec, c.validators[fakeType][defaultScope])
}

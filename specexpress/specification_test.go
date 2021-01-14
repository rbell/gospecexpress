// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/rbell/gospecexpress/interfaces/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSpecification_ForType_ShouldReturnValidatorBuilder(t *testing.T) {
	// setup
	type testSubject struct {
		Name string
	}
	ts := &testSubject{}
	spec := &Specification{
		validators: &sync.Map{},
	}

	// test
	result := spec.ForType(ts)

	// assert
	assert.Equal(t, reflect.TypeOf(ts), spec.forType)
	assert.NotNil(t, result)
}

func TestSpecification_Validate_ShouldCallValidator(t *testing.T) {
	// setup
	type testSubject struct {
		Name string
	}
	ts := &testSubject{}
	mValidator := &mocks.Validator{}
	mValidator.On("Validate", ts, map[string]interface{}(nil), mock.Anything).Return(nil)

	spec := &Specification{
		validators: &sync.Map{},
		forType:    reflect.TypeOf(ts),
	}

	addValidator(spec.validators, "Name", "Alias0", mValidator)

	// test
	err := spec.Validate(ts, nil)

	// assert
	mValidator.AssertExpectations(t)
	assert.Nil(t, err)
}

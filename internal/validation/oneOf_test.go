// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/rbell/gospecexpress/interfaces/mocks"
)

func TestOneOf_Validate_MatchesValue_ShouldReturnNil(t *testing.T) {
	// setup
	validator := &OneOf{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "Enum",
		},
		values: []interface{}{"test1", "test2"},
	}

	mMessageStore := &mocks.MessageStorer{}

	type testSubjectType struct {
		Enum string
	}
	testSubject := &testSubjectType{Enum: "test2"}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
	mMessageStore.AssertExpectations(t)
}

func TestOneOf_Validate_DoesNotMatchValue_ShouldReturnError(t *testing.T) {
	// setup
	validator := &OneOf{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "Enum",
		},
		values: []interface{}{"test1", "test2"},
	}

	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.OneOf"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Does Not Match")

	type testSubjectType struct {
		Enum string
	}
	testSubject := &testSubjectType{Enum: "test3"}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

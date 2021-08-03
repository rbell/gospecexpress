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

func TestOneOfFold_Validate_MatchesValue_ShouldReturnNil(t *testing.T) {
	// setup
	validator := &OneOfFold{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "Enum",
		},
		values: []string{"test1", "test2"},
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

func TestOneOfFold_Validate_DoesNotMatchValue_ShouldReturnError(t *testing.T) {
	// setup
	validator := &OneOfFold{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "Enum",
		},
		values: []string{"test1", "test2"},
	}

	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.OneOfFold"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Does Not Match")

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

func TestOneOfFold_Validate_ReferenceMatchValue_ShouldReturnNil(t *testing.T) {
	// setup
	validator := &OneOfFold{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "Enum",
		},
		values: []string{"test1", "test2"},
	}

	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.OneOfFold"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Does Not Match")

	type testSubjectType struct {
		Enum *string
	}
	test := "test2"
	testSubject := &testSubjectType{Enum: &test}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestOneOfFold_Validate_ReferenceDoesNotMatchValue_ShouldReturnErr(t *testing.T) {
	// setup
	validator := &OneOfFold{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "Enum",
		},
		values: []string{"test1", "test2"},
	}

	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.OneOfFold"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Does Not Match")

	type testSubjectType struct {
		Enum *string
	}
	test := "test3"
	testSubject := &testSubjectType{Enum: &test}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

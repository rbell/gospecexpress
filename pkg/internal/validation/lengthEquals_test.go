// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestLengthEquals_Validate_LengthsEqual_ShouldReturnNil(t *testing.T) {
	// setup
	validator := &LengthEquals{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		lenEq: 2,
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		TestField []string
	}
	testSubject := &testSubjectType{TestField: []string{
		"test1",
		"test2",
	}}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestLengthEquals_Validate_LengthLessThanRequiredLen_ShouldReturnNil(t *testing.T) {
	// setup
	validator := &LengthEquals{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		lenEq: 5,
	}
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LengthEquals"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Length Not Equal")

	type testSubjectType struct {
		TestField []string
	}
	testSubject := &testSubjectType{TestField: []string{
		"test1",
		"test2",
	}}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

func TestLengthEquals_Validate_LengthGreaterThanRequiredLen_ShouldReturnNil(t *testing.T) {
	// setup
	validator := &LengthEquals{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		lenEq: 1,
	}
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LengthEquals"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Length Not Equal")

	type testSubjectType struct {
		TestField []string
	}
	testSubject := &testSubjectType{TestField: []string{
		"test1",
		"test2",
	}}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

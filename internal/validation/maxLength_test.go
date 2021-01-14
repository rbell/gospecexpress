// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"gitlab.com/rbell/gospecexpress/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

func TestMaxLength_Validate_ShouldReturnNilWhenLengthLessThanMaxLength(t *testing.T) {
	// setup
	validator := &MaxLength{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "FirstName",
		},
		maxLen: 50,
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fred"}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestMaxLength_Validate_ShouldReturnErrorWhenLengthGreaterThanMaxLength(t *testing.T) {
	// setup
	validator := &MaxLength{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "FirstName",
		}, maxLen: 5,
	}
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.MaxLength"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Length Not Less Than")
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fred Flinstone"}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

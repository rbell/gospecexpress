// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/rbell/gospecexpress/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

func TestMinLength_Validate_ShouldReturnNilWhenLengthGreaterThanMinLength(t *testing.T) {
	// setup
	validator := &MinLength{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "FirstName",
		},
		minLen: 5,
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fredrick"}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestMinLength_Validate_ShouldReturnErrorWhenLengthLessThanMinLength(t *testing.T) {
	// setup
	validator := &MinLength{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "FirstName",
		}, minLen: 5,
	}
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.MinLength"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Length Not Greater Than")
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fred"}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

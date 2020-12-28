package validation

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"

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

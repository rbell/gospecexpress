package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

func TestMaxLength_Validate_ShouldReturnNilWhenLengthLessThanMaxLength(t *testing.T) {
	// setup
	validator := &MaxLength{
		AllFieldValidators: &AllFieldValidators{
			FieldName: "FirstName",
		},
		maxLen: 50,
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fred"}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestMaxLength_Validate_ShouldReturnErrorWhenLengthGreaterThanMaxLength(t *testing.T) {
	// setup
	validator := &MaxLength{
		AllFieldValidators: &AllFieldValidators{
			FieldName: "FirstName",
		}, maxLen: 5,
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fred Flinstone"}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

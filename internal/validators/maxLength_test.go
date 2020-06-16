package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLength_Validate_ShouldReturnNilWhenLengthLessThanMaxLength(t *testing.T) {
	// setup
	validator := &MaxLength{
		fieldName: "FirstName",
		maxLen:    50,
	}
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fred"}

	// test
	result := validator.Validate(testSubject)

	// assert
	assert.Nil(t, result)
}

func TestMaxLength_Validate_ShouldReturnErrorWhenLengthGreaterThanMaxLength(t *testing.T) {
	// setup
	validator := &MaxLength{
		fieldName: "FirstName",
		maxLen:    5,
	}
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fred Flinstone"}

	// test
	result := validator.Validate(testSubject)

	// assert
	assert.NotNil(t, result)
}

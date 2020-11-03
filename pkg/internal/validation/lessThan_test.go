package validation

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestLessThan_Validate_ValueIsLessThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := LessThanValue("Distance", 50)
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 40}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestLessThan_Validate_ValueIsEqual_ShouldReturnErr(t *testing.T) {
	// setup
	validator := LessThanValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThan")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 40}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestLessThan_Validate_ValueIsGreater_ShouldReturnErr(t *testing.T) {
	// setup
	validator := LessThanValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThan")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 50}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

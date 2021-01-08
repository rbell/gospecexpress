package validation

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestBetweenValues_Validate_FieldValueBetween_ShouldReturnNil(t *testing.T) {
	// setup
	validator := BetweenValues("Distance", "Distance", 50, 100)
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 75}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
	mMessageStore.AssertExpectations(t)
}

func TestBetweenValues_Validate_FieldValueLessThanLower_ShouldReturnError(t *testing.T) {
	// setup
	validator := BetweenValues("Distance", "Distance", 50, 100)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.Between"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not Between")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 25}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

func TestBetweenValues_Validate_FieldValueGreaterThanUpper_ShouldReturnError(t *testing.T) {
	// setup
	validator := BetweenValues("Distance", "Distance", 50, 100)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.Between"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not Between")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 125}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

func TestBetweenOtherFieldValues_Validate_FieldValueBetween_ShouldReturnNil(t *testing.T) {
	// setup
	validator := BetweenOtherFieldValues("Distance", "Distance", "Lower", "Upper")
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int
		Lower    int
		Upper    int
	}
	testSubject := &testSubjectType{Distance: 75, Lower: 50, Upper: 100}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
	mMessageStore.AssertExpectations(t)
}

func TestBetweenOtherFieldValues_Validate_FieldValueLessThanLower_ShouldReturnError(t *testing.T) {
	// setup
	validator := BetweenOtherFieldValues("Distance", "Distance", "Lower", "Upper")
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.Between"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not Between")
	type testSubjectType struct {
		Distance int
		Lower    int
		Upper    int
	}
	testSubject := &testSubjectType{Distance: 25, Lower: 50, Upper: 100}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

func TestBetweenOtherFieldValues_Validate_FieldValueGreaterThanUpper_ShouldReturnError(t *testing.T) {
	// setup
	validator := BetweenOtherFieldValues("Distance", "Distance", "Lower", "Upper")
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.Between"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not Between")
	type testSubjectType struct {
		Distance int
		Lower    int
		Upper    int
	}
	testSubject := &testSubjectType{Distance: 125, Lower: 50, Upper: 100}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestLessThanOrEqual_Validate_ValueIsLessThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := LessThanOrEqualToValue("Distance", 50)
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

func TestLessThanOrEqual_Validate_ValueIsEqual_ShouldReturnNil(t *testing.T) {
	// setup
	validator := LessThanOrEqualToValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThanEqual")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 40}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestLessThanOrEqual_Validate_ValueIsGreater_ShouldReturnErr(t *testing.T) {
	// setup
	validator := LessThanOrEqualToValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThanEqual")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 50}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestLessThanOrEqualToValueFromContext_Validate_ValueIsLessThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := LessThanOrEqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 40,
		DistanceB: 50,
	}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestLessThanOrEqualToValueFromContext_Validate_ValueIsEqualTo_ShouldReturnNil(t *testing.T) {
	// setup
	validator := LessThanOrEqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 50,
		DistanceB: 50,
	}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestLessThanOrEqualToValueFromContext_Validate_ValueIsGreaterThan_ShouldReturnErr(t *testing.T) {
	// setup
	validator := LessThanOrEqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThan")
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 60,
		DistanceB: 50,
	}

	// test
	result := validator.Validate(testSubject, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

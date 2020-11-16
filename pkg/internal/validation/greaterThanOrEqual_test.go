package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestGreaterThanOrEqualTo_Validate_ValueIsGreaterThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := GreaterThanOrEqualToValue("Distance", 50)
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 60}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestGreaterThanOrEqualTo_Validate_ValueIsEqual_ShouldReturnNil(t *testing.T) {
	// setup
	validator := GreaterThanOrEqualToValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.GreaterThanThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not GreaterThanThanEqual")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 40}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestGreaterThanOrEqualTo_Validate_ValueIsLess_ShouldReturnErr(t *testing.T) {
	// setup
	validator := GreaterThanOrEqualToValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.GreaterThanEqual"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not GreaterThanThanEqual")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 30}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestGreaterThanThanOrEqualToValueFromContext_Validate_ValueIsGreaterThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := GreaterThanOrEqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 40,
		DistanceB: 30,
	}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestGreaterThanOrEqualToValueFromContext_Validate_ValueIsEqualTo_ShouldReturnNil(t *testing.T) {
	// setup
	validator := GreaterThanOrEqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
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
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestGreaterThanThanOrEqualToValueFromContext_Validate_ValueIsLessThan_ShouldReturnErr(t *testing.T) {
	// setup
	validator := GreaterThanOrEqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.GreaterThanEqual"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not GreaterThanThan")
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 40,
		DistanceB: 50,
	}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

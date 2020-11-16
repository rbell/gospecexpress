package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestEqualTo_Validate_ValueIsEqualTo_ShouldReturnNil(t *testing.T) {
	// setup
	validator := EqualToValue("Distance", 50)
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 50}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestEqualTo_Validate_ValueIsGreaterThan_ShouldReturnErr(t *testing.T) {
	// setup
	validator := EqualToValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.EqualTo"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not EqualTo")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 50}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestEqualTo_Validate_ValueIsLess_ShouldReturnErr(t *testing.T) {
	// setup
	validator := EqualToValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.EqualTo"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not EqualTo")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 30}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestEqualToValueFromContext_Validate_ValueIsEqual_ShouldReturnNil(t *testing.T) {
	// setup
	validator := EqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
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

func TestEqualToValueFromContext_Validate_ValueIsGreaterThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := EqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.EqualTo"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not EqualTo")
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 60,
		DistanceB: 50,
	}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestEqualToValueFromContext_Validate_ValueIsLessThan_ShouldReturnErr(t *testing.T) {
	// setup
	validator := EqualToValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.EqualTo"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not EqualTo")
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
package validation

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestContains_Validate_StringContainsValue_ShouldReturnNil(t *testing.T) {
	// Setup
	type testSubjectType struct {
		TestField string
	}
	testSubj := &testSubjectType{TestField: "This is a test"}
	contains := 's'

	validator := &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		contains:    contains,
		containsVal: reflect.ValueOf(contains),
	}
	mMsgStore := &mocks.MessageStorer{}

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.Nil(t, result)
	mMsgStore.AssertExpectations(t)
}

func TestContains_Validate_StringDoesNotContainValue_ShouldReturnError(t *testing.T) {
	// Setup
	type testSubjectType struct {
		TestField string
	}
	testSubj := &testSubjectType{TestField: "This is a test"}
	contains := 'z'

	validator := &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		contains:    contains,
		containsVal: reflect.ValueOf(contains),
	}
	mMsgStore := &mocks.MessageStorer{}
	mMsgStore.On("GetMessage", mock.AnythingOfType("*validation.Contains"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not in slice")

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.NotNil(t, result)
	mMsgStore.AssertExpectations(t)
}

func TestContains_Validate_SliceContainsValue_ShouldReturnError(t *testing.T) {
	// Setup
	type testElement struct {
		Name string
	}
	type testSubjectType struct {
		TestField []testElement
	}

	testSubj := &testSubjectType{TestField: []testElement{
		{Name: "test1"},
		{Name: "test2"},
		{Name: "test3"},
	}}
	contains := testSubj.TestField[2]

	validator := &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		contains:    contains,
		containsVal: reflect.ValueOf(contains),
	}
	mMsgStore := &mocks.MessageStorer{}

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.Nil(t, result)
	mMsgStore.AssertExpectations(t)
}

func TestContains_Validate_SliceDoesNotContainValue_ShouldReturnError(t *testing.T) {
	// Setup
	type testElement struct {
		Name string
	}
	type testSubjectType struct {
		TestField []testElement
	}

	testSubj := &testSubjectType{TestField: []testElement{
		{Name: "test1"},
		{Name: "test2"},
		{Name: "test3"},
	}}
	contains := testElement{Name: "test4"}

	validator := &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		contains:    contains,
		containsVal: reflect.ValueOf(contains),
	}
	mMsgStore := &mocks.MessageStorer{}
	mMsgStore.On("GetMessage", mock.AnythingOfType("*validation.Contains"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not in slice")

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.NotNil(t, result)
	mMsgStore.AssertExpectations(t)
}

func TestContains_Validate_SliceOfRefsContainsValue_ShouldReturnError(t *testing.T) {
	// Setup
	type testElement struct {
		Name string
	}
	type testSubjectType struct {
		TestField []*testElement
	}

	testSubj := &testSubjectType{TestField: []*testElement{
		{Name: "test1"},
		{Name: "test2"},
		{Name: "test3"},
	}}
	contains := testSubj.TestField[2]

	validator := &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		contains:    contains,
		containsVal: reflect.ValueOf(contains),
	}
	mMsgStore := &mocks.MessageStorer{}

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.Nil(t, result)
	mMsgStore.AssertExpectations(t)
}

func TestContains_Validate_SliceOfRefsDoesNotContainValue_ShouldReturnError(t *testing.T) {
	// Setup
	type testElement struct {
		Name string
	}
	type testSubjectType struct {
		TestField []*testElement
	}

	testSubj := &testSubjectType{TestField: []*testElement{
		{Name: "test1"},
		{Name: "test2"},
		{Name: "test3"},
	}}
	contains := &testElement{Name: "test4"}

	validator := &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		contains:    contains,
		containsVal: reflect.ValueOf(contains),
	}
	mMsgStore := &mocks.MessageStorer{}
	mMsgStore.On("GetMessage", mock.AnythingOfType("*validation.Contains"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not in slice")

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.NotNil(t, result)
	mMsgStore.AssertExpectations(t)
}

func TestContains_Validate_SliceTypeDoesNotMatchValueType_ShouldReturnError(t *testing.T) {
	// Setup
	type testElement struct {
		Name string
	}
	type testSubjectType struct {
		TestField []*testElement
	}

	testSubj := &testSubjectType{TestField: []*testElement{
		{Name: "test1"},
		{Name: "test2"},
		{Name: "test3"},
	}}
	contains := "this is not the same type as testElement"

	validator := &Contains{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		contains:    contains,
		containsVal: reflect.ValueOf(contains),
	}
	mMsgStore := &mocks.MessageStorer{}
	mMsgStore.On("GetMessage", mock.AnythingOfType("*validation.Contains"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not in slice")

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.NotNil(t, result)
	mMsgStore.AssertExpectations(t)
}

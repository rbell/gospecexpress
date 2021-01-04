package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/errors"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestRangeValidate_Validate_ValidSlice_ShouldReturnNil(t *testing.T) {
	// Setup
	type testElement struct {
		Name string
	}
	type testSubjectType struct {
		TestField []testElement
	}
	testSubj := &testSubjectType{
		TestField: []testElement{
			{"This is a test"},
		},
	}

	// Mock call to cataloger.ValidateWithContext for reference returning nil (valid)
	mCataloger := &mocks.Cataloger{}
	mCataloger.On("ValidateWithContext", testSubj.TestField[0], map[string]interface{}(nil)).Return(nil)

	validator := &RangeValidate{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		validationCatalog: mCataloger,
	}
	mMsgStore := &mocks.MessageStorer{}

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.Nil(t, result)
	mMsgStore.AssertExpectations(t)
	mCataloger.AssertExpectations(t)
}

func TestRangeValidate_Validate_InvalidSliceElementZero_ShouldReturnError(t *testing.T) {
	// Setup
	type testElement struct {
		Name string
	}
	type testSubjectType struct {
		TestField []testElement
	}
	testSubj := &testSubjectType{
		TestField: []testElement{
			{"This is a test"},
		},
	}

	// Mock call to cataloger.ValidateWithContext for reference returning nil (valid)
	mCataloger := &mocks.Cataloger{}
	mCataloger.On("ValidateWithContext", testSubj.TestField[0], map[string]interface{}(nil)).Return(errors.NewValidationError("Name", "Invalid"))

	validator := &RangeValidate{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		validationCatalog: mCataloger,
	}
	mMsgStore := &mocks.MessageStorer{}

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.NotNil(t, result)
	//nolint:errcheck // ignore error cause we are testing
	ve, _ := result.(*errors.ValidatorError)
	assert.Contains(t, ve.GetFlatErrorMap(), "TestField[0].Name")
	mMsgStore.AssertExpectations(t)
	mCataloger.AssertExpectations(t)
}

func TestRangeValidate_Validate_InvalidSliceElementOne_ShouldReturnError(t *testing.T) {
	// Setup
	type testElement struct {
		Name string
	}
	type testSubjectType struct {
		TestField []testElement
	}
	testSubj := &testSubjectType{
		TestField: []testElement{
			{"This is a test1"},
			{"This is a test2"},
		},
	}

	// Mock call to cataloger.ValidateWithContext for reference returning nil (valid)
	mCataloger := &mocks.Cataloger{}
	mCataloger.On("ValidateWithContext", testSubj.TestField[0], map[string]interface{}(nil)).Return(nil)
	mCataloger.On("ValidateWithContext", testSubj.TestField[1], map[string]interface{}(nil)).Return(errors.NewValidationError("Name", "Invalid"))

	validator := &RangeValidate{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		validationCatalog: mCataloger,
	}
	mMsgStore := &mocks.MessageStorer{}

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.NotNil(t, result)
	//nolint:errcheck // ignore error cause we are testing
	ve, _ := result.(*errors.ValidatorError)
	assert.Contains(t, ve.GetFlatErrorMap(), "TestField[1].Name")
	mMsgStore.AssertExpectations(t)
	mCataloger.AssertExpectations(t)
}

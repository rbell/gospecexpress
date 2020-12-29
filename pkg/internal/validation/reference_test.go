package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestReference_Validate_ValidReference_ShouldReturnNil(t *testing.T) {
	// Setup
	type TestReference struct {
		Adddress string
		City     string
	}
	type TestSubject struct {
		TestRef *TestReference
	}

	testSubj := &TestSubject{
		TestRef: &TestReference{
			Adddress: "123 West Way",
			City:     "Somewhere",
		},
	}

	// Mock call to cataloger.ValidateWithContext for reference returning nil (valid)
	mCataloger := &mocks.Cataloger{}
	mCataloger.On("ValidateWithContext", testSubj.TestRef, map[string]interface{}(nil)).Return(nil)

	refValidator := &Reference{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestRef",
		},
		validationCatalog: mCataloger,
	}

	// Test
	result := refValidator.Validate(testSubj, nil, nil)

	// Assert
	assert.Nil(t, result)
	mCataloger.AssertExpectations(t)
}

func TestReference_Validate_InValidReference_ShouldReturnNil(t *testing.T) {
	// Setup
	type TestReference struct {
		Adddress string
		City     string
	}
	type TestSubject struct {
		TestRef *TestReference
	}

	testSubj := &TestSubject{
		TestRef: &TestReference{
			Adddress: "123 West Way",
			City:     "Somewhere",
		},
	}

	// Mock call to cataloger.ValidateWithContext for reference returning nil (valid)
	mCataloger := &mocks.Cataloger{}
	mCataloger.On("ValidateWithContext", testSubj.TestRef, map[string]interface{}(nil)).Return(NewValidationError("TestRef", "Invalid"))

	refValidator := &Reference{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestRef",
		},
		validationCatalog: mCataloger,
	}

	// Test
	result := refValidator.Validate(testSubj, nil, nil)

	// Assert
	assert.NotNil(t, result)
	mCataloger.AssertExpectations(t)
	valErr, ok := result.(*ValidatorError)
	assert.True(t, ok)
	assert.Len(t, valErr.errorMap, 1)
	assert.Contains(t, valErr.errorMap, "TestRef")
}

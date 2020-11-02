package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

func TestValidate_ValidForPopulatedRequiredField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		AllFieldValidators: &AllFieldValidators{
			FieldName: "FirstName",
		},
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		FirstName string
	}
	testSubject := &testSubjectType{FirstName: "Fred"}

	// Test
	result := validator.Validate(testSubject, mMessageStore)

	// Assert
	assert.Nil(t, result)
}

func TestValidate_NonExportedField_ValidForPopulatedRequiredField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		AllFieldValidators: &AllFieldValidators{
			FieldName: "firstName",
		},
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		firstName string
	}
	testSubject := &testSubjectType{firstName: "Fred"}

	// Test
	result := validator.Validate(testSubject, mMessageStore)

	// Assert
	assert.Nil(t, result)
}

func TestValidate_NotValidForUnPopulatedRequiredField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		AllFieldValidators: &AllFieldValidators{
			FieldName: "firstName",
		},
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		firstName string
	}
	testSubject := &testSubjectType{firstName: ""}

	// Test
	result := validator.Validate(testSubject, mMessageStore)

	// Assert
	assert.NotNil(t, result)
}

func TestValidate_Numeric_ValidForNonZeroField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		AllFieldValidators: &AllFieldValidators{
			FieldName: "Distance",
		},
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int64
	}
	testSubject := &testSubjectType{Distance: int64(100)}

	// Test
	result := validator.Validate(testSubject, mMessageStore)

	// Assert
	assert.Nil(t, result)
}

func TestValidate_Numeric_NotValidForZeroField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		AllFieldValidators: &AllFieldValidators{
			FieldName: "Distance",
		},
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int64
	}
	testSubject := &testSubjectType{Distance: int64(0)}

	// Test
	result := validator.Validate(testSubject, mMessageStore)

	// Assert
	assert.NotNil(t, result)
}

// TODO: Pointers (nil references)?
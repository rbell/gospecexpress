package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_ValidForPopulatedRequiredField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		fieldName:   "FirstName",
	}
	type testSubectType struct {
		FirstName string
	}
	testSubject := &testSubectType{FirstName: "Fred"}

	// Test
	result := validator.Validate(testSubject)

	// Assert
	assert.True(t, result)
}

func TestValidate_NonExportedField_ValidForPopulatedRequiredField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		fieldName:   "firstName",
	}
	type testSubectType struct {
		firstName string
	}
	testSubject := &testSubectType{firstName: "Fred"}

	// Test
	result := validator.Validate(testSubject)

	// Assert
	assert.True(t, result)
}

func TestValidate_NotValidForUnPopulatedRequiredField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		fieldName:   "",
	}
	type testSubectType struct {
		firstName string
	}
	testSubject := &testSubectType{firstName: "Fred"}

	// Test
	result := validator.Validate(testSubject)

	// Assert
	assert.False(t, result)
}

func TestValidate_Numeric_ValidForNonZeroField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		fieldName:   "Distance",
	}
	type testSubectType struct {
		Distance int64
	}
	testSubject := &testSubectType{Distance: int64(100)}

	// Test
	result := validator.Validate(testSubject)

	// Assert
	assert.True(t, result)
}

func TestValidate_Numeric_NotValidForZeroField(t *testing.T) {
	// Setup
	validator := &RequiredField{
		fieldName:   "Distance",
	}
	type testSubectType struct {
		Distance int64
	}
	testSubject := &testSubectType{Distance: int64(0)}

	// Test
	result := validator.Validate(testSubject)

	// Assert
	assert.False(t, result)
}

// TODO: Pointers (nil references)?

package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/errors"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestRangeExect_Validate_ValidSlice_ShouldReturnNil(t *testing.T) {
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

	validator := &RangeExpect{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		exp: func(ctx interfaces.ValidatorContextGetter) error {
			return nil
		},
	}
	mMsgStore := &mocks.MessageStorer{}

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.Nil(t, result)
	mMsgStore.AssertExpectations(t)
}

func TestRangeExect_Validate_InvalidValidSlice_ShouldReturnNil(t *testing.T) {
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

	validator := &RangeExpect{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "TestField",
		},
		exp: func(ctx interfaces.ValidatorContextGetter) error {
			return errors.NewValidationError("TestField", "Invalid")
		},
	}
	mMsgStore := &mocks.MessageStorer{}

	// Test
	result := validator.Validate(testSubj, nil, mMsgStore)

	// Assert
	assert.NotNil(t, result)
	//nolint:errcheck // ignore error cause we are testing
	ve, _ := result.(*errors.ValidatorError)
	assert.Contains(t, ve.GetFlatErrorMap(), "TestField[0].TestField")
	mMsgStore.AssertExpectations(t)
}

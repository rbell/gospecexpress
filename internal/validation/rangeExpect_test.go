// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"testing"

	"github.com/rbell/gospecexpress/errors"

	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/interfaces/mocks"
	"github.com/stretchr/testify/assert"
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
		exp: func(ctx interfaces.FieldValidatorContextGetter) error {
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
		exp: func(ctx interfaces.FieldValidatorContextGetter) error {
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

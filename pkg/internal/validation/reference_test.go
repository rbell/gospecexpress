// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/errors"

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

func TestReference_Validate_InValidReference_ShouldReturnError(t *testing.T) {
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
	mCataloger.On("ValidateWithContext", testSubj.TestRef, map[string]interface{}(nil)).Return(errors.NewValidationError("Address", "Invalid"))

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
	valErr, ok := result.(*errors.ValidatorError)
	assert.True(t, ok)
	errs := valErr.GetFlatErrorMap()
	assert.Len(t, errs, 1)
	assert.Contains(t, errs, "TestRef.Address")
}

// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"testing"

	"github.com/rbell/gospecexpress/interfaces"
	"github.com/rbell/gospecexpress/interfaces/mocks"
	"github.com/stretchr/testify/assert"
)

func TestExpectation_Validate_Expectations_Pass_ShouldReturn_Nil(t *testing.T) {
	// Setup
	validator := &Expectation{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "Country",
		},
		exp: func(thing interface{}, ctx interfaces.ValidatorContextGetter) error {
			// Fake business logic where only US or CA valid countries (i.e. valid shipping countries)
			cntry := ctx.GetFieldValue("Country")
			if !(cntry == "US" || cntry == "CA") {
				return fmt.Errorf("Invalid country %v", cntry)
			}
			return nil
		},
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Country string
	}
	testSubject := &testSubjectType{Country: "US"}

	// Test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// Assert
	assert.Nil(t, result)
}

func TestExpectation_Validate_Expectations_DoesNotPass_ShouldReturn_error(t *testing.T) {
	// Setup
	validator := &Expectation{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "Country",
		},
		exp: func(thing interface{}, ctx interfaces.ValidatorContextGetter) error {
			// Fake business logic where only US or CA valid countries (i.e. valid shipping countries)
			cntry := ctx.GetFieldValue("Country")
			if !(cntry == "US" || cntry == "CA") {
				return fmt.Errorf("Invalid country %v", cntry)
			}
			return nil
		},
	}
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Country string
	}
	testSubject := &testSubjectType{Country: "UK"}

	// Test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "Invalid country UK\n", result.Error())

}

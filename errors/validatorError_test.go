// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorError_GetErrorMap_NoHierarchy_ReturnsFlatMap(t *testing.T) {
	// Setup
	ve := &ValidatorError{
		errorMap: map[string][]string{
			"TestField": {
				"Err1",
				"Err2",
			},
		},
	}

	// Test
	result := ve.GetFlatErrorMap()

	// Assert
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Contains(t, result, "TestField")
	assert.Len(t, result["TestField"], 2)
}

func TestValidatorError_GetErrorMap_OneChild_ReturnsFlatMap(t *testing.T) {
	// Setup
	ve := &ValidatorError{
		errorMap: map[string][]string{
			"TestField": {
				"Err1",
				"Err2",
			},
		},
		children: map[string]*ValidatorError{
			"TestRef": {
				errorMap: map[string][]string{
					"ChildTestField": {
						"ChildErr1",
						"ChildErr2",
						"ChildErr3",
					},
				},
			},
		},
	}

	// Test
	result := ve.GetFlatErrorMap()

	// Assert
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Contains(t, result, "TestField")
	assert.Len(t, result["TestField"], 2)
	assert.Contains(t, result, "TestRef.ChildTestField")
	assert.Len(t, result["TestRef.ChildTestField"], 3)
}

func TestValidatorError_GetErrorMap_OneGrandChild_ReturnsFlatMap(t *testing.T) {
	// Setup
	ve := &ValidatorError{
		errorMap: map[string][]string{
			"TestField": {
				"Err1",
				"Err2",
			},
		},
		children: map[string]*ValidatorError{
			"TestRef": {
				errorMap: map[string][]string{
					"ChildTestField": {
						"ChildErr1",
						"ChildErr2",
						"ChildErr3",
					},
				},
				children: map[string]*ValidatorError{
					"GrandChildTestRef": {
						errorMap: map[string][]string{
							"GrandChildTestField": {
								"GrandChildErr1",
								"GrandChildErr2",
								"GrandChildErr3",
								"GrandChildErr4",
							},
						},
					},
				},
			},
		},
	}

	// Test
	result := ve.GetFlatErrorMap()

	// Assert
	assert.NotNil(t, result)
	assert.Len(t, result, 3)
	assert.Contains(t, result, "TestField")
	assert.Len(t, result["TestField"], 2)
	assert.Contains(t, result, "TestRef.ChildTestField")
	assert.Len(t, result["TestRef.ChildTestField"], 3)
	assert.Contains(t, result, "TestRef.GrandChildTestRef.GrandChildTestField")
	assert.Len(t, result["TestRef.GrandChildTestRef.GrandChildTestField"], 4)
}

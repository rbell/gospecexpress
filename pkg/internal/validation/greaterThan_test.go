// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestGreaterThan_Validate_ValueIsGreaterThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := GreaterThanValue("Distance", 50)
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 60}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestGreaterThan_Validate_ValueIsEqual_ShouldReturnErr(t *testing.T) {
	// setup
	validator := GreaterThanValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.GreaterThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not GreaterThan")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 40}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestGreaterThan_Validate_ValueIsLess_ShouldReturnErr(t *testing.T) {
	// setup
	validator := GreaterThanValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.GreaterThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not GreaterThan")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 30}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestGreaterThanValueFromContext_Validate_ValueIsGreaterThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := GreaterThanValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 60,
		DistanceB: 50,
	}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestGreaterThanValueFromContext_Validate_ValueIsEqualTo_ShouldReturnErr(t *testing.T) {
	// setup
	validator := GreaterThanValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.GreaterThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not GreaterThan")
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 50,
		DistanceB: 50,
	}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestGreaterThanValueFromContext_Validate_ValueIsGreaterThan_ShouldReturnErr(t *testing.T) {
	// setup
	validator := GreaterThanValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.GreaterThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not GreaterThan")
	type testSubjectType struct {
		DistanceA int
		DistanceB int
	}
	testSubject := &testSubjectType{
		DistanceA: 40,
		DistanceB: 50,
	}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

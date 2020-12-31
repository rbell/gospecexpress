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

func TestLessThan_Validate_ValueIsLessThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := LessThanValue("Distance", 50)
	mMessageStore := &mocks.MessageStorer{}
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 40}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
}

func TestLessThan_Validate_ValueIsEqual_ShouldReturnErr(t *testing.T) {
	// setup
	validator := LessThanValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThan")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 40}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestLessThan_Validate_ValueIsGreater_ShouldReturnErr(t *testing.T) {
	// setup
	validator := LessThanValue("Distance", 40)
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThan")
	type testSubjectType struct {
		Distance int
	}
	testSubject := &testSubjectType{Distance: 50}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
}

func TestLessThanValueFromContext_Validate_ValueIsLessThan_ShouldReturnNil(t *testing.T) {
	// setup
	validator := LessThanValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
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
	assert.Nil(t, result)
}

func TestLessThanValueFromContext_Validate_ValueIsEqualTo_ShouldReturnErr(t *testing.T) {
	// setup
	validator := LessThanValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThan")
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

func TestLessThanValueFromContext_Validate_ValueIsGreaterThan_ShouldReturnErr(t *testing.T) {
	// setup
	validator := LessThanValueFromContext("DistanceA", func(ctx interfaces.ValidatorContextGetter) interface{} {
		return ctx.GetFieldValue("DistanceB")
	})
	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.LessThan"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Not LessThan")
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
	assert.NotNil(t, result)
}

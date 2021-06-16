// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	errors "github.com/rbell/gospecexpress/errors"
	interfaces "github.com/rbell/gospecexpress/interfaces"

	mock "github.com/stretchr/testify/mock"
)

// ValidationExpression is an autogenerated mock type for the ValidationExpression type
type ValidationExpression struct {
	mock.Mock
}

// Execute provides a mock function with given fields: thing, ctx
func (_m *ValidationExpression) Execute(thing interface{}, ctx interfaces.ValidatorContextGetter) *errors.ValidatorError {
	ret := _m.Called(thing, ctx)

	var r0 *errors.ValidatorError
	if rf, ok := ret.Get(0).(func(interface{}, interfaces.ValidatorContextGetter) *errors.ValidatorError); ok {
		r0 = rf(thing, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ValidatorError)
		}
	}

	return r0
}

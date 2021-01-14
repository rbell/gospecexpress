// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	interfaces "gitlab.com/rbell/gospecexpress/interfaces"
)

// MessageStorer is an autogenerated mock type for the MessageStorer type
type MessageStorer struct {
	mock.Mock
}

// GetMessage provides a mock function with given fields: validator, ctx
func (_m *MessageStorer) GetMessage(validator interfaces.Validator, ctx interfaces.ValidatorContextGetter) string {
	ret := _m.Called(validator, ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(interfaces.Validator, interfaces.ValidatorContextGetter) string); ok {
		r0 = rf(validator, ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetMessage provides a mock function with given fields: validator, getterFunc
func (_m *MessageStorer) SetMessage(validator interfaces.Validator, getterFunc interfaces.ErrorMessageGetterFunc) {
	_m.Called(validator, getterFunc)
}
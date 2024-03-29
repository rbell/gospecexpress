// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/rbell/gospecexpress/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// MessageOverrider is an autogenerated mock type for the MessageOverrider type
type MessageOverrider struct {
	mock.Mock
}

// GetOverrideErrorMessage provides a mock function with given fields: ctx
func (_m *MessageOverrider) GetOverrideErrorMessage(ctx interfaces.FieldValidatorContextGetter) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(interfaces.FieldValidatorContextGetter) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetOverrideErrorMessage provides a mock function with given fields: msgFormatter
func (_m *MessageOverrider) SetOverrideErrorMessage(msgFormatter interfaces.MessageFormatter) {
	_m.Called(msgFormatter)
}

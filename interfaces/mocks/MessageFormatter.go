// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/rbell/gospecexpress/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// MessageFormatter is an autogenerated mock type for the MessageFormatter type
type MessageFormatter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx
func (_m *MessageFormatter) Execute(ctx interfaces.ValidatorContextGetter) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(interfaces.ValidatorContextGetter) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

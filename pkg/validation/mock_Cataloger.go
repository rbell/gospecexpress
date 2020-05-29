// Code generated by mockery v1.0.0. DO NOT EDIT.

package validation

import mock "github.com/stretchr/testify/mock"

// MockCataloger is an autogenerated mock type for the Cataloger type
type MockCataloger struct {
	mock.Mock
}

// RegisterForType provides a mock function with given fields: forType, v
func (_m *MockCataloger) RegisterForType(forType interface{}, v Specification) {
	_m.Called(forType, v)
}

// Validate provides a mock function with given fields: something
func (_m *MockCataloger) Validate(something interface{}) bool {
	ret := _m.Called(something)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(something)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

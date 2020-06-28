// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	interfaces "gitlab.com/govalidate/pkg/interfaces"

	reflect "reflect"
)

// SpecificationValidator is an autogenerated mock type for the SpecificationValidator type
type SpecificationValidator struct {
	mock.Mock
}

// ForType provides a mock function with given fields: forType
func (_m *SpecificationValidator) ForType(forType interface{}) interfaces.QualifierBuilder {
	ret := _m.Called(forType)

	var r0 interfaces.QualifierBuilder
	if rf, ok := ret.Get(0).(func(interface{}) interfaces.QualifierBuilder); ok {
		r0 = rf(forType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.QualifierBuilder)
		}
	}

	return r0
}

// GetForType provides a mock function with given fields:
func (_m *SpecificationValidator) GetForType() reflect.Type {
	ret := _m.Called()

	var r0 reflect.Type
	if rf, ok := ret.Get(0).(func() reflect.Type); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.Type)
		}
	}

	return r0
}

// Validate provides a mock function with given fields: _a0
func (_m *SpecificationValidator) Validate(_a0 interface{}) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

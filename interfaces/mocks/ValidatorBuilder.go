// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/rbell/gospecexpress/interfaces"
	mock "github.com/stretchr/testify/mock"

	regexp "regexp"
)

// ValidatorBuilder is an autogenerated mock type for the ValidatorBuilder type
type ValidatorBuilder struct {
	mock.Mock
}

// Between provides a mock function with given fields: lower, upper, options
func (_m *ValidatorBuilder) Between(lower interface{}, upper interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, lower, upper)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(lower, upper, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// BetweenOtherFields provides a mock function with given fields: lowerField, upperField, options
func (_m *ValidatorBuilder) BetweenOtherFields(lowerField string, upperField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, lowerField, upperField)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(string, string, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(lowerField, upperField, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// BetweenValuesFromContext provides a mock function with given fields: lowerGetter, upperGetter, options
func (_m *ValidatorBuilder) BetweenValuesFromContext(lowerGetter interfaces.ValueFromContext, upperGetter interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, lowerGetter, upperGetter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.ValueFromContext, interfaces.ValueFromContext, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(lowerGetter, upperGetter, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// Contains provides a mock function with given fields: thing, options
func (_m *ValidatorBuilder) Contains(thing interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, thing)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interface{}, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(thing, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// ContainsValueFromContext provides a mock function with given fields: fromContext, options
func (_m *ValidatorBuilder) ContainsValueFromContext(fromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromContext)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.ValueFromContext, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(fromContext, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// EqualTo provides a mock function with given fields: value, options
func (_m *ValidatorBuilder) EqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interface{}, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(value, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// EqualToOtherField provides a mock function with given fields: otherField, options
func (_m *ValidatorBuilder) EqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, otherField)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(string, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(otherField, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// EqualToValueFromContext provides a mock function with given fields: valueFromContext, options
func (_m *ValidatorBuilder) EqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, valueFromContext)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.ValueFromContext, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(valueFromContext, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// Expect provides a mock function with given fields: validator, options
func (_m *ValidatorBuilder) Expect(validator interfaces.FieldValidationExpression, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, validator)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.FieldValidationExpression, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(validator, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// GreaterThan provides a mock function with given fields: value, options
func (_m *ValidatorBuilder) GreaterThan(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interface{}, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(value, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// GreaterThanOrEqualTo provides a mock function with given fields: value, options
func (_m *ValidatorBuilder) GreaterThanOrEqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interface{}, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(value, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// GreaterThanOrEqualToOtherField provides a mock function with given fields: otherField, options
func (_m *ValidatorBuilder) GreaterThanOrEqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, otherField)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(string, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(otherField, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// GreaterThanOrEqualToValueFromContext provides a mock function with given fields: valueFromContext, options
func (_m *ValidatorBuilder) GreaterThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, valueFromContext)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.ValueFromContext, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(valueFromContext, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// GreaterThanOtherField provides a mock function with given fields: otherField, options
func (_m *ValidatorBuilder) GreaterThanOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, otherField)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(string, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(otherField, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// GreaterThanValueFromContext provides a mock function with given fields: valueFromContext, options
func (_m *ValidatorBuilder) GreaterThanValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, valueFromContext)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.ValueFromContext, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(valueFromContext, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// If provides a mock function with given fields: condition
func (_m *ValidatorBuilder) If(condition interfaces.FieldValidationCondition) interfaces.ValidatorBuilder {
	ret := _m.Called(condition)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.FieldValidationCondition) interfaces.ValidatorBuilder); ok {
		r0 = rf(condition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// LengthEquals provides a mock function with given fields: length, options
func (_m *ValidatorBuilder) LengthEquals(length int, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, length)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(int, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(length, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// LessThan provides a mock function with given fields: value, options
func (_m *ValidatorBuilder) LessThan(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interface{}, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(value, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// LessThanOrEqualTo provides a mock function with given fields: value, options
func (_m *ValidatorBuilder) LessThanOrEqualTo(value interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interface{}, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(value, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// LessThanOrEqualToOtherField provides a mock function with given fields: otherField, options
func (_m *ValidatorBuilder) LessThanOrEqualToOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, otherField)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(string, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(otherField, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// LessThanOrEqualToValueFromContext provides a mock function with given fields: valueFromContext, options
func (_m *ValidatorBuilder) LessThanOrEqualToValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, valueFromContext)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.ValueFromContext, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(valueFromContext, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// LessThanOtherField provides a mock function with given fields: otherField, options
func (_m *ValidatorBuilder) LessThanOtherField(otherField string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, otherField)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(string, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(otherField, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// LessThanValueFromContext provides a mock function with given fields: valueFromContext, options
func (_m *ValidatorBuilder) LessThanValueFromContext(valueFromContext interfaces.ValueFromContext, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, valueFromContext)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(interfaces.ValueFromContext, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(valueFromContext, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// Matches provides a mock function with given fields: regex, regexDescripton, options
func (_m *ValidatorBuilder) Matches(regex *regexp.Regexp, regexDescripton string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, regex, regexDescripton)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(*regexp.Regexp, string, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(regex, regexDescripton, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// MaxLength provides a mock function with given fields: len, options
func (_m *ValidatorBuilder) MaxLength(length int, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, len)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(int, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(length, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// MinLength provides a mock function with given fields: len, options
func (_m *ValidatorBuilder) MinLength(length int, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, len)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(int, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(length, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// OneOf provides a mock function with given fields: values, options
func (_m *ValidatorBuilder) OneOf(values []interface{}, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, values)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func([]interface{}, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(values, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// Optional provides a mock function with given fields: fieldName
func (_m *ValidatorBuilder) Optional(fieldName string) interfaces.ValidatorBuilder {
	ret := _m.Called(fieldName)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(string) interfaces.ValidatorBuilder); ok {
		r0 = rf(fieldName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// RangeExpect provides a mock function with given fields: validator, options
func (_m *ValidatorBuilder) RangeExpect(validator func(interfaces.FieldValidatorContextGetter) error, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, validator)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(func(interfaces.FieldValidatorContextGetter) error, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(validator, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// RangeValidate provides a mock function with given fields: options
func (_m *ValidatorBuilder) RangeValidate(options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// Required provides a mock function with given fields: fieldName, options
func (_m *ValidatorBuilder) Required(fieldName string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fieldName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(string, ...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(fieldName, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// ValidateReference provides a mock function with given fields: options
func (_m *ValidatorBuilder) ValidateReference(options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interfaces.ValidatorBuilder
	if rf, ok := ret.Get(0).(func(...interfaces.ValidatorOption) interfaces.ValidatorBuilder); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ValidatorBuilder)
		}
	}

	return r0
}

// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"strings"
)

// ValidatorError contains any validation errors.  Each error is associated with a some validation context (i.e. field name error is for, etc)
type ValidatorError struct {
	errorMap   map[string][]string
	warningMap map[string][]string
	// children contains errors for fields that reference another validated structure
	children map[string]*ValidatorError
}

// NewValidationError returns a new validation error
func NewValidationError(context, msg string, isWarning bool) *ValidatorError {
	em := make(map[string][]string)
	em[context] = []string{msg}
	if isWarning {
		return &ValidatorError{warningMap: em, errorMap: make(map[string][]string)}
	}
	return &ValidatorError{errorMap: em, warningMap: make(map[string][]string)}
}

// NewValidationErrors returns a new validation error for a map of error messages
func NewValidationErrors(errs map[string][]string, children map[string]*ValidatorError) *ValidatorError {
	if errs == nil {
		return &ValidatorError{errorMap: make(map[string][]string), children: children}
	}
	return &ValidatorError{errorMap: errs, children: children}
}

// IsValidatorError returns reference to ValidatorError and a bool indicating if the err passed in is a ValidatorError
func IsValidatorError(err error) (*ValidatorError, bool) {
	if e, ok := err.(*ValidatorError); ok {
		return e, true
	}
	return nil, false
}

// Error returns the error messages in a single string
func (e *ValidatorError) Error() string {
	sb := strings.Builder{}
	for _, ee := range e.GetFlatErrorMap() {
		for _, e := range ee {
			sb.WriteString(fmt.Sprintf("ERROR: %v\n", e))
		}
	}
	for _, ee := range e.GetFlatWarningMap() {
		for _, e := range ee {
			sb.WriteString(fmt.Sprintf("WARNING: %v\n", e))
		}
	}
	return sb.String()
}

// GetFlatErrorMap gets a map of error messages mapped by field
func (e *ValidatorError) GetFlatErrorMap() map[string][]string {
	flatMap := e.errorMap
	for k, v := range e.children {
		childErrors := getFlattenedMap(k, v, false)
		for ek, e := range childErrors {
			addMsgs(flatMap, ek, e...)
		}
	}
	return flatMap
}

// GetFlatWarningMap gets a map of warning messages mapped by field
func (e *ValidatorError) GetFlatWarningMap() map[string][]string {
	flatMap := e.warningMap
	for k, v := range e.children {
		childErrors := getFlattenedMap(k, v, true)
		for ek, e := range childErrors {
			addMsgs(flatMap, ek, e...)
		}
	}
	return flatMap
}

// GetErrorMap returns the top level error messages
func (e *ValidatorError) GetErrorMap() map[string][]string {
	return e.errorMap
}

// GetWarningMap returns the top level warning messages
func (e *ValidatorError) GetWarningMap() map[string][]string {
	return e.warningMap
}

// GetChildErrors returns errors of structs referenced by the struct being validated
func (e *ValidatorError) GetChildErrors() map[string]*ValidatorError {
	return e.children
}

func getFlattenedMap(key string, ve *ValidatorError, getWarnings bool) map[string][]string {
	flatMap := prefixKeys(ve.errorMap, key+".")
	if getWarnings {
		flatMap = prefixKeys(ve.warningMap, key+".")
	}
	for childKey, child := range ve.children {
		// resursively call geFlattenedMap
		childMap := getFlattenedMap(key+"."+childKey, child, getWarnings)
		for k, e := range childMap {
			addMsgs(flatMap, k, e...)
		}
	}
	return flatMap
}

func prefixKeys(m map[string][]string, prefix string) map[string][]string {
	result := make(map[string][]string)
	for k, v := range m {
		result[prefix+k] = v
	}
	return result
}

func addMsgs(errMap map[string][]string, context string, msg ...string) {
	if _, ok := errMap[context]; !ok {
		errMap[context] = []string{}
	}
	errMap[context] = append(errMap[context], msg...)
}

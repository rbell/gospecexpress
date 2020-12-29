package errors

import (
	"fmt"
	"strings"
)

// ValidatorError contains any validation errors.  Each error is associated with a some validation context (i.e. field name error is for, etc)
type ValidatorError struct {
	errorMap map[string][]string
	// children contains errors for fields that reference another validated structure
	children map[string]*ValidatorError
}

// NewValidationError returns a new validation error
func NewValidationError(context, msg string) *ValidatorError {
	em := make(map[string][]string)
	em[context] = []string{msg}
	return &ValidatorError{errorMap: em}
}

// NewValidationErrors returns a new validation error for a map of error messages
func NewValidationErrors(errs map[string][]string, children map[string]*ValidatorError) *ValidatorError {
	if errs == nil {
		return &ValidatorError{errorMap: make(map[string][]string), children: children}
	}
	return &ValidatorError{errorMap: errs, children: children}
}

// Error returns the error messages in a single string
func (e *ValidatorError) Error() string {
	sb := strings.Builder{}
	for _, ee := range e.GetFlatErrorMap() {
		for _, e := range ee {
			sb.WriteString(fmt.Sprintf("%v\n", e))
		}
	}
	return sb.String()
}

// GetFlatErrorMap gets a map of error messages mapped by field
func (e *ValidatorError) GetFlatErrorMap() map[string][]string {
	flatMap := e.errorMap
	for k, v := range e.children {
		childErrors := getFlattenedMap(k, v)
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

// GetChildErrors returns errors of structs referenced by the struct being validated
func (e *ValidatorError) GetChildErrors() map[string]*ValidatorError {
	return e.children
}

func getFlattenedMap(key string, ve *ValidatorError) map[string][]string {
	flatMap := prefixKeys(ve.errorMap, key+".")
	for childKey, child := range ve.children {
		// resursively call geFlattenedMap
		childMap := getFlattenedMap(key+"."+childKey, child)
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

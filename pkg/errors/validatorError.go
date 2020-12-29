package errors

import (
	"fmt"
	"strings"
)

// ValidatorError contains any validation errors.  Each error is associated with a some validation context (i.e. field name error is for, etc)
type ValidatorError struct {
	errorMap map[string][]string
}

// NewValidationError returns a new validation error
func NewValidationError(context, msg string) *ValidatorError {
	em := make(map[string][]string)
	em[context] = []string{msg}
	return &ValidatorError{errorMap: em}
}

// NewValidationErrors returns a new validation error for a map of error messages
func NewValidationErrors(errs map[string][]string) *ValidatorError {
	return &ValidatorError{errorMap: errs}
}

// Error returns the error messages in a single string
func (e *ValidatorError) Error() string {
	sb := strings.Builder{}
	for _, ee := range e.errorMap {
		for _, e := range ee {
			sb.WriteString(fmt.Sprintf("%v\n", e))
		}
	}
	return sb.String()
}

// GetErrorMap gets a map of error messages mapped by field
func (e *ValidatorError) GetErrorMap() map[string][]string {
	return e.errorMap
}

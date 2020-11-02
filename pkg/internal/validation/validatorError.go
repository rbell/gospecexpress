package validation

import (
	"errors"
	"fmt"
	"reflect"
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

// JoinErrors joins two errors together into a ValidatorError.
func JoinErrors(e1, e2 error) *ValidatorError {
	var e *ValidatorError
	if (e1 == nil || reflect.ValueOf(e1).IsNil()) && e2 != nil {
		if errors.As(e2, &e) {
			return e2.(*ValidatorError)
		}
		return NewValidationError("", e2.Error())

	}

	var ve *ValidatorError
	if errors.As(e1, &e) {
		//nolint:errcheck // above line infers its castable
		ve = e1.(*ValidatorError)
	} else {
		ve = NewValidationError("", e1.Error())
	}

	if errors.As(e2, &e) {
		for key, msg := range e2.(*ValidatorError).errorMap {
			ve = ve.AddMsgs(key, msg...)
		}
	}

	return ve
}

// AddMsgs adds an error message to the error
func (e *ValidatorError) AddMsgs(context string, msg ...string) *ValidatorError {
	if _, ok := e.errorMap[context]; !ok {
		e.errorMap[context] = []string{}
	}
	e.errorMap[context] = append(e.errorMap[context], msg...)

	return e
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

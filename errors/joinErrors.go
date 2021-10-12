// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"reflect"
)

// JoinErrors joins two errors together into a ValidatorError
func JoinErrors(e1, e2 error) *ValidatorError {
	var e *ValidatorError
	// if e1 is nil and e2 is not, short circuit using e2
	if (e1 == nil || reflect.ValueOf(e1).IsNil()) && e2 != nil {
		if errors.As(e2, &e) {
			// e1 is nil and e2 is a ValidationError, return e2
			return e2.(*ValidatorError)
		}
		// e1 is nil and e2 is an error.  Return new validation error using
		return NewValidationError("", e2.Error(), false)
	}

	var ve *ValidatorError
	if errors.As(e1, &e) {
		//nolint:errcheck // above line infers its castable
		ve = e1.(*ValidatorError)
	} else {
		ve = NewValidationError("", e1.Error(), false)
	}

	if errors.As(e2, &e) {
		errMap := ve.GetErrorMap()
		for key, msg := range e2.(*ValidatorError).GetFlatErrorMap() {
			AddMessagesToMap(errMap, key, msg...)
		}
		warnMap := ve.GetWarningMap()
		for key, msg := range e2.(*ValidatorError).GetFlatWarningMap() {
			AddMessagesToMap(warnMap, key, msg...)
		}
		childErrs := ve.GetChildErrors()
		for key, ve := range e2.(*ValidatorError).GetChildErrors() {
			childErrs[key] = ve
		}
	}

	return ve
}

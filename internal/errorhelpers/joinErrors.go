// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errorhelpers

import (
	"errors"
	"reflect"

	specExpressErrors "github.com/rbell/gospecexpress/errors"
)

// JoinErrors joins two errors together into a ValidatorError
func JoinErrors(e1, e2 error) *specExpressErrors.ValidatorError {
	var e *specExpressErrors.ValidatorError
	// if e1 is nil and e2 is not, short circuit using e2
	if (e1 == nil || reflect.ValueOf(e1).IsNil()) && e2 != nil {
		if errors.As(e2, &e) {
			// e1 is nil and e2 is a ValidationError, return e2
			return e2.(*specExpressErrors.ValidatorError)
		}
		// e1 is nil and e2 is an error.  Return new validation error using
		return specExpressErrors.NewValidationError("", e2.Error(), false)
	}

	var ve *specExpressErrors.ValidatorError
	if errors.As(e1, &e) {
		//nolint:errcheck // above line infers its castable
		ve = e1.(*specExpressErrors.ValidatorError)
	} else {
		ve = specExpressErrors.NewValidationError("", e1.Error(), false)
	}

	if errors.As(e2, &e) {
		errMap := ve.GetErrorMap()
		for key, msg := range e2.(*specExpressErrors.ValidatorError).GetFlatErrorMap() {
			AddMessages(errMap, key, msg...)
		}
		warnMap := ve.GetWarningMap()
		for key, msg := range e2.(*specExpressErrors.ValidatorError).GetFlatWarningMap() {
			AddMessages(warnMap, key, msg...)
		}
		childErrs := ve.GetChildErrors()
		for key, ve := range e2.(*specExpressErrors.ValidatorError).GetChildErrors() {
			childErrs[key] = ve
		}
	}

	return ve
}

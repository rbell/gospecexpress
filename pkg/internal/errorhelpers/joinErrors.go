// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errorhelpers

import (
	"errors"
	"reflect"

	specExpressErrors "gitlab.com/rbell/gospecexpress/pkg/errors"
)

// JoinErrors joins two errors together into a ValidatorError
func JoinErrors(e1, e2 error) *specExpressErrors.ValidatorError {
	var e *specExpressErrors.ValidatorError
	if (e1 == nil || reflect.ValueOf(e1).IsNil()) && e2 != nil {
		if errors.As(e2, &e) {
			return e2.(*specExpressErrors.ValidatorError)
		}
		return specExpressErrors.NewValidationError("", e2.Error())

	}

	var ve *specExpressErrors.ValidatorError
	if errors.As(e1, &e) {
		//nolint:errcheck // above line infers its castable
		ve = e1.(*specExpressErrors.ValidatorError)
	} else {
		ve = specExpressErrors.NewValidationError("", e1.Error())
	}

	if errors.As(e2, &e) {
		errMap := ve.GetErrorMap()
		for key, msg := range e2.(*specExpressErrors.ValidatorError).GetFlatErrorMap() {
			AddMessages(errMap, key, msg...)
		}
		childErrs := ve.GetChildErrors()
		for key, ve := range e2.(*specExpressErrors.ValidatorError).GetChildErrors() {
			childErrs[key] = ve
		}
	}

	return ve
}

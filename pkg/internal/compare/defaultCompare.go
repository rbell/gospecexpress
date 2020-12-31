// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compare

import (
	"errors"
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"
)

// errUnknownComparisonResult is error returned when comparison cannot be determined
var errUnknownComparisonResult = errors.New("unable to compare")

// DefaultComparer defines a comparer which attempts to implement compare if the instance does not implement its own
type DefaultComparer struct {
	aVal reflect.Value
}

// NewDefaultComparer returns an initialized DefaultComparer
func NewDefaultComparer(a interface{}) *DefaultComparer {
	return &DefaultComparer{
		aVal: reflect.ValueOf(a),
	}
}

// Compare compares A to B, returning  -1 if A < B, 0 if A == B, 1 if A > B, error if A cannot be compared to B
func (c *DefaultComparer) Compare(b interface{}) (int, error) {
	bVal := reflect.ValueOf(b)
	if lt, e := reflectionhelpers.Lt(c.aVal, bVal); e != nil {
		return 0, e
	} else if lt {
		return -1, nil
	}

	if eq, e := reflectionhelpers.Eq(c.aVal, bVal); e != nil {
		return 0, e
	} else if eq {
		return 0, nil
	}

	if gt, e := reflectionhelpers.Gt(c.aVal, bVal); e != nil {
		return 0, e
	} else if gt {
		return 1, nil
	}

	return 0, errUnknownComparisonResult
}

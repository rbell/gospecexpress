// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testmodels

import "time"

// Customer is a sample customer model for purposes of validation
type Customer struct {
	FirstName      string
	middleName     string
	LastName       string
	Age            int
	MemberSince    time.Time
	MemberExpireAt time.Time
}

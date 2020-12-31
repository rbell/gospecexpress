// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/examples/simpletest/testmodels"
	// import specifications, but not referenced.  Need to do so to execute init methods defined in the package
	_ "gitlab.com/rbell/gospecexpress/examples/simpletest/testspec"
)

func main() {
	// We have something we need to validate: a customer
	c := &testmodels.Customer{
		FirstName:      "",
		LastName:       "Flinstone",
		Age:            23,
		MemberSince:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		MemberExpireAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// Validate the instance c against specifications in the catalog
	err := catalog.ValidationCatalog().Validate(c)

	if err == nil {
		fmt.Printf("Customer is valid.")
	} else {
		// error contains messages as to what is invalid
		fmt.Printf("Customer is not valid:\n%v", err.Error())
	}
}

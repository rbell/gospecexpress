# Go SpecExpress
Go SpecExpress is a fluent validation library for Go that makes it easy to consolidate the aspect of validation into specifications.

# Features
- Consolidate set of validation rules for a given type in a specification, separating the concern of validation from the object being validated.
- Specifications registered in Catalog of specifications allowing a simple call to validate a value or reference as simple as `err := catalog.Validate(instance)`
- Wide variety of validators including support for different types including slices, maps, structs, etc.
- Support comparing data in same type to each other (i.e. `Required("EndDate").GreaterThanOtherField("StartDate"))`)
- Support for validating data against data outside the type being validated 
- Support for validation and error reporting on complex composite structs (i.e. Order.ShipAddress)
- Customizable error messaging

# Requirements
Go 1.14 or above

# Installation
Run the following to install the package:
```
go get gitlab.com/rbell/gospecexpress
```

# Quick Start

```go
// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	. "gitlab.com/rbell/gospecexpress/pkg/specexpress"
)

// ClubMember is a sample customer model for purposes of validation
type ClubMember struct {
	FirstName      string
	middleName     string
	LastName       string
	Age            int
	MemberSince    time.Time
	MemberExpireAt time.Time
}

// ClubMemberSpec defines a specification for a customer
type ClubMemberSpec struct {
	Specification
}

func newClubMemberSpec() *ClubMemberSpec {
	s := &ClubMemberSpec{}

	s.ForType(&ClubMember{}).
		RequiredField("FirstName", OverrideMessage("The First Name is a required field!")).MaxLength(5).
		Required("LastName").MaxLength(50).
		Required("Age").LessThan(80).
		Required("MemberExpireAt").GreaterThanOtherField("MemberSince")

	return s
}

func init() {
	// Register the ClubMemberSpec in the Catalog.
	// Registering on init ensures the catalog is initialized, however, registration can happen anytime before the catalog is used.
	catalog.Register(newClubMemberSpec())
}

func main() {
	// We have something we need to validate: a customer
	c := &ClubMember{
		FirstName:      "",
		LastName:       "Flinstone",
		Age:            23,
		MemberSince:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		MemberExpireAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// Validate the instance c against specifications in the catalog
	err := catalog.Validate(c)

	if err == nil {
		fmt.Printf("ClubMember is valid.")
	} else {
		// error contains messages as to what is invalid
		fmt.Printf("ClubMember is not valid:\n%v", err.Error())
	}
}
```
Output:
```
ClubMember is not valid:
The First Name is a required field!
MemberExpireAt should be Greater than 2020-01-01 00:00:00 +0000 UTC.
```


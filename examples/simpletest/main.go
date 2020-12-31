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
		Required("FirstName", OverrideMessage("The First Name is a required field!")).MaxLength(5).
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

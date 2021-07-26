// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/rbell/gospecexpress"

	"github.com/rbell/gospecexpress/catalog"
)

// ClubMember is a sample customer model for purposes of validation
// Field names can be tagged with a user readable field name used when referencing the field in error messaging
type ClubMember struct {
	FirstName      string `spec:"First Name"`
	MiddleName     string
	LastName       string
	CreditScore    int
	Age            int
	MemberSince    time.Time
	MemberExpireAt time.Time
	Guardian       string
}

// ClubMemberSpec defines a specification for a customer
type ClubMemberSpec struct {
	gospecexpress.Specification
}

func newClubDefaultMemberSpec() *ClubMemberSpec {
	s := &ClubMemberSpec{}

	// Validation does not include validation of Guardian
	s.ForType(&ClubMember{}).
		Required("FirstName").MaxLength(50).
		Optional("MiddleName").MaxLength(20).
		Required("LastName", gospecexpress.WithErrorMessage("Sir Name is a required field!")).MaxLength(50).
		Required("CreditScore").GreaterThan(640, gospecexpress.WithWarning()).
		Required("MemberExpireAt").GreaterThanOtherField("MemberSince")

	return s
}

func newClubMemberSpecForAMinor() *ClubMemberSpec {
	s := &ClubMemberSpec{}

	// Validation includes validation of Guardian
	s.ForType(&ClubMember{}).
		// Adding scope signalling this is a specification for a Minor (i.e. someone under age of 18)
		ForScope("MinorClubMember", true).
		Required("Guardian").If(func(t interface{}, c map[string]interface{}) bool { return isMinor(t, c) }).MaxLength(10)

	return s
}

func init() {
	// Register the ClubMemberSpec in the Catalog.
	// Registering on init ensures the catalog is initialized, however, registration can happen anytime before the catalog is used.
	catalog.Register(newClubDefaultMemberSpec())
	catalog.Register(newClubMemberSpecForAMinor())
}

func main() {
	// We have something we need to validate: a customer
	c := &ClubMember{
		FirstName:      "Fred",
		LastName:       "Flinstone",
		Age:            16,
		CreditScore:    700,
		Guardian:       "",
		MemberSince:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		MemberExpireAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// Validate the instance c using default specification registered in the catalog
	err := catalog.Validate(c)

	if err == nil {
		fmt.Printf("ClubMember is valid using the default specification.\n")
	} else {
		// error contains messages as to what is invalid
		fmt.Printf("ClubMember is not valid using the default specification:\n%v", err.Error())
	}

	// Validate the instance using the "MinorClubMember" specification registered in the catalog
	err = catalog.Validate(c, catalog.WithScope("MinorClubMember"))

	if err == nil {
		fmt.Printf("ClubMember is valid using the MinorClubMember specification.\n")
	} else {
		// error contains messages as to what is invalid
		fmt.Printf("ClubMember is not valid using the MinorClubMember specification:\n%v", err.Error())
	}
}

func isMinor(thing interface{}, contextData map[string]interface{}) bool {
	if cm, ok := thing.(*ClubMember); ok {
		return cm.Age < 18
	}
	return false
}

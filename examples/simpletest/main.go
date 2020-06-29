package main

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/specificationcatalog"

	"gitlab.com/rbell/gospecexpress/examples/simpletest/testmodels"
	// import specifications, but not referenced.  Need to do so to execute init methods defined in the package
	_ "gitlab.com/rbell/gospecexpress/examples/simpletest/testspec"
)

func main() {
	// We have something we need to validate: a customer
	c := &testmodels.Customer{
		FirstName: "Fred Flinstone",
		LastName:  "Flinstone",
		Age:       23,
	}

	// Validate it against the specifications we have registered in the specification catalog
	// (specification registers itself via init function in testspec/customerSpec.go)
	err := specificationcatalog.Catalog().Validate(c)
	if err == nil {
		fmt.Printf("Customer is valid.")
	} else {
		// error contains messages as to what is invalid
		fmt.Printf("Customer is not valid:\n%v", err.Error())
	}
}

package main

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/catalog"

	"gitlab.com/rbell/gospecexpress/examples/simpletest/testmodels"
	// import specifications, but not referenced.  Need to do so to execute init methods defined in the package
	_ "gitlab.com/rbell/gospecexpress/examples/simpletest/testspec"
)

func main() {
	// Example of overriding a default message for a specific validator
	//specificationcatalog.ValidationCatalog().MessageStore().SetMessage(&validators.MaxLength{}, func(ctx *errors.ErrorMessageContext) string {
	//	return "Too Long!!!"
	//})

	// We have something we need to validate: a customer
	c := &testmodels.Customer{
		FirstName: "",
		LastName:  "Flinstone",
		Country:   "UK",
		Age:       23,
		DistanceA: 40,
		DistanceB: 30,
		Handicap:  98,
	}

	// Validate it against the specifications we have registered in the specification catalog
	// (specification registers itself via init function in testspec/customerSpec.go)
	err := catalog.ValidationCatalog().ValidateWithContext(c, map[string]interface{}{
		"MaximumHandicap": 80,
	})

	if err == nil {
		fmt.Printf("Customer is valid.")
	} else {
		// error contains messages as to what is invalid
		fmt.Printf("Customer is not valid:\n%v", err.Error())
	}
}

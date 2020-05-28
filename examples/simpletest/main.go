package main

import (
	"fmt"

	"local/gospecexpress/examples/simpletest/testmodels"
	"local/gospecexpress/pkg/validation"

	_ "local/gospecexpress/examples/simpletest/testspec"
)

func main() {
	c := &testmodels.Customer{
		FirstName: "Fred",
		LastName:  "Flinstone",
	}
	isvalid := validation.Catalog().Validate(c)
	fmt.Printf("Customer valid: %v", isvalid)
}


package main

import (
	"fmt"
	"gitlab.com/govalidate/examples/simpletest/testmodels"
	"gitlab.com/govalidate/pkg/specexpress"

	_ "gitlab.com/govalidate/examples/simpletest/testspec"
)

func main() {
	c := &testmodels.Customer{
		FirstName: "Fred",
		LastName:  "Flinstone",
	}
	isvalid := specexpress.Catalog().Validate(c)
	fmt.Printf("Customer valid: %v", isvalid)
}


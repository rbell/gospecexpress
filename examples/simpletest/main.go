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
		LastName:  "",
		Age: 23,
	}
	err := specexpress.Catalog().Validate(c)
	if err == nil {
		fmt.Printf("Customer is valid.")
	} else {
		fmt.Printf("Customer is not valid:\n%v", err.Error())
	}
}


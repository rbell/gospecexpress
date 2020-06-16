package testspec

import (
	"gitlab.com/govalidate/examples/simpletest/testmodels"
	"gitlab.com/govalidate/pkg/specexpress"
)

// init functions run at first import, registering the specification in the specification catalog
// (we can define multiple init functions in the same package and they all will get executed upon import)
func init() {
	specexpress.Catalog().Register(newTestSpec())
}

// CustomerSpec defines a specification for a customer
type CustomerSpec struct {
	specexpress.Specification
}

func newTestSpec() *CustomerSpec {
	s := &CustomerSpec{}

	s.ForType(&testmodels.Customer{}).
		RequiredField("FirstName").MaxLength(5).
		RequiredField("LastName").MaxLength(50)

	return s
}

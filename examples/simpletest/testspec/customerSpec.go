package testspec

import (
	"gitlab.com/govalidate/examples/simpletest/testmodels"
	"gitlab.com/govalidate/pkg/specexpress"
)

func init() {
	// init runs at first import, registering the specification in the specification catalog
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

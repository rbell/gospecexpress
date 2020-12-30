package testspec

import (
	"gitlab.com/rbell/gospecexpress/examples/simpletest/testmodels"
	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	. "gitlab.com/rbell/gospecexpress/pkg/specexpress"
)

func init() {
	// Register the CustomerSpec in the Catalog.
	// Registering at init ensures the catalog is initialized, however, registration can happen anytime before the catalog is used.
	catalog.ValidationCatalog().Register(newTestSpec())
}

// CustomerSpec defines a specification for a customer
type CustomerSpec struct {
	Specification
}

func newTestSpec() *CustomerSpec {
	s := &CustomerSpec{}

	s.ForType(&testmodels.Customer{}).
		RequiredField("FirstName", OverrideMessage("The First Name is a required field!")).MaxLength(5).
		Required("LastName").MaxLength(50).
		Required("Age").LessThan(80).
		Required("MemberExpireAt").GreaterThanOtherField("MemberSince")

	return s
}

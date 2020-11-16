package testspec

import (
	"gitlab.com/rbell/gospecexpress/examples/simpletest/testmodels"
	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/specexpress"
)

// init functions run at first import, registering the specification in the specification catalog
// (we can define multiple init functions in the same package and they all will get executed upon import)
func init() {
	catalog.ValidationCatalog().Register(newTestSpec())
}

// CustomerSpec defines a specification for a customer
type CustomerSpec struct {
	specexpress.Specification
}

func newTestSpec() *CustomerSpec {
	s := &CustomerSpec{}

	s.ForType(&testmodels.Customer{}).
		RequiredField("FirstName").MaxLength(5).
		RequiredField("LastName").MaxLength(50).
		RequiredField("Age").LessThan(80).
		RequiredField("DistanceA").LessThanOtherField("DistanceB").
		RequiredField("Handicap").LessThanValueFromContext(
		func(ctx interfaces.ValidatorContextGetter) interface{} {
			// If a MaximumHandicap was passed in, then make sure customer handicap less than that, otherwise default to 100
			data := ctx.GetContextData()
			if maxHandicap, ok := data["MaximumHandicap"]; ok {
				return maxHandicap
			}
			return 100
		})

	return s
}

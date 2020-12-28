package testspec

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/examples/simpletest/testmodels"
	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	. "gitlab.com/rbell/gospecexpress/pkg/specexpress"
)

// init functions run at first import, registering the specification in the specification catalog
// (we can define multiple init functions in the same package and they all will get executed upon import)
func init() {
	catalog.ValidationCatalog().Register(newTestSpec())
}

// Fake some cache of valid countries (i.e. valid ship to countries)
var validCountriesCache = []string{"US", "CA"}

// CustomerSpec defines a specification for a customer
type CustomerSpec struct {
	Specification
}

func newTestSpec() *CustomerSpec {
	s := &CustomerSpec{}

	s.ForType(&testmodels.Customer{}).
		RequiredField("FirstName", OverrideMessage("The First Name is a required field!")).MaxLength(5).
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
		}).
		RequiredField("Country").
		Expect(
			func(ctx interfaces.ValidatorContextGetter) error {
				// Fake business logic where only US or CA valid countries (i.e. valid shipping countries)
				cntry := ctx.GetFieldValue("Country")
				for _, c := range validCountriesCache {
					if c == cntry {
						return nil
					}
				}
				return fmt.Errorf("Invalid country %v", cntry)
			})

	return s
}

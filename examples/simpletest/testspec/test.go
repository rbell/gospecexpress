package testspec

import (
	"gitlab.com/govalidate/examples/simpletest/testmodels"
	"gitlab.com/govalidate/pkg/specexpress"
)

func init() {
	specexpress.Catalog().Register(newTestSpec())

}

type TestSpec struct {
	specexpress.Specification
}

func newTestSpec() *TestSpec {
	s := &TestSpec{}

	s.ForType(&testmodels.Customer{}).
		RequiredField("FirstName").MaxLength(5).
		RequiredField("LastName").MaxLength(50)

	return s
}

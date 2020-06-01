package testspec

import (
	"gitlab.com/govalidate/examples/simpletest/testmodels"
	"gitlab.com/govalidate/pkg/specexpress"
)

func init() {
	specexpress.Catalog().Register(newTestSpec().ForType(&testmodels.Customer{}))
}

var _ specexpress.SpecificationBuilder = &TestSpec{}

type TestSpec struct {
	specexpress.Specification
}

func newTestSpec() *TestSpec {
	return &TestSpec{}
}

func (t *TestSpec) Validate(thing interface{}) bool {
	return true
}


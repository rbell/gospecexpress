package testspec

import (
	"local/gospecexpress/examples/simpletest/testmodels"
	"local/gospecexpress/pkg/validation"
)

func init() {
	validation.Catalog().RegisterForType(&testmodels.Customer{}, &TestSpec{})
}

var _ validation.Validator = &TestSpec{}

type TestSpec struct {

}

func (t *TestSpec) IsValid() bool {
	return true
}

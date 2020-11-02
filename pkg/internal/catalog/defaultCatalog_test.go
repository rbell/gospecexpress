package catalog

import (
	"reflect"
	"testing"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCatalog_RegisterForType_ShouldRegisterDefalutSpecForType(t *testing.T) {
	// setup
	c := &DefaultCatalog{validators: make(map[reflect.Type]map[string]interfaces.SpecificationValidator)}
	mSpec := &mocks.SpecificationValidator{}
	type fakeStruct struct{}
	fake := &fakeStruct{}
	fakeType := reflect.TypeOf(fake)

	mSpec.On("getForType").Return(fakeType)

	// test
	c.Register(mSpec)

	// assert
	mSpec.AssertExpectations(t)
	assert.Contains(t, c.validators, fakeType)
	assert.Contains(t, c.validators[fakeType], defaultContext)
	assert.Equal(t, mSpec, c.validators[fakeType][defaultContext])
}

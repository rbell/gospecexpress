package specexpress

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/govalidate/pkg/interfaces"
	"gitlab.com/govalidate/pkg/interfaces/mocks"
)

func TestCatalog_RegisterForType_ShouldRegisterDefalutSpecForType(t *testing.T) {
	// setup
	c := &catalog{validators: make(map[reflect.Type]map[string]interfaces.SpecificationValidator)}
	mSpec := &mocks.SpecificationValidator{}
	type fakeStruct struct{}
	fake := &fakeStruct{}
	fakeType := reflect.TypeOf(fake)

	mSpec.On("GetForType").Return(fakeType)

	// test
	c.Register(mSpec)

	// assert
	mSpec.AssertExpectations(t)
	assert.Contains(t, c.validators, fakeType)
	assert.Contains(t, c.validators[fakeType], defaultContext)
	assert.Equal(t, mSpec, c.validators[fakeType][defaultContext])
}

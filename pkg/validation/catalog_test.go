package validation

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCatalog_RegisterForType_ShouldRegisterDefalutSpecForType(t *testing.T) {
	// setup
	c := &catalog{validators: make(map[reflect.Type]map[string]Specification)}
	mSpec := &MockSpecification{}
	type fakeStruct struct {}
	fake := &fakeStruct{}
	fakeType := reflect.TypeOf(fake)

	// test
	c.RegisterForType(fake, mSpec)

	// assert
	assert.Contains(t, c.validators, fakeType)
	assert.Contains(t, c.validators[fakeType], defaultContext)
	assert.Equal(t, mSpec, c.validators[fakeType][defaultContext])
}

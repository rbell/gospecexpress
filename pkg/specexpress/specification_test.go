package specexpress

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/govalidate/pkg/interfaces"
	"gitlab.com/govalidate/pkg/interfaces/mocks"
)

func TestSpecification_ForType_ShouldReturnValidatorBuilder(t *testing.T) {
	// setup
	type testSubject struct {
		Name string
	}
	ts := &testSubject{}
	spec := &Specification{
		validators: []interfaces.Validator{},
	}

	// test
	result := spec.ForType(ts)

	// assert
	assert.Equal(t, reflect.TypeOf(ts), spec.forType)
	assert.NotNil(t, result)
}

func TestSpecification_Validate_ShouldCallValidator(t *testing.T) {
	// setup
	type testSubject struct {
		Name string
	}
	ts := &testSubject{}
	mValidator := &mocks.Validator{}
	mValidator.On("Validate", ts).Return(nil)

	spec := &Specification{
		validators: []interfaces.Validator{mValidator},
		forType:    reflect.TypeOf(ts),
	}

	// test
	err := spec.Validate(ts)

	// assert
	mValidator.AssertExpectations(t)
	assert.Nil(t, err)
}

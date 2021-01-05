package validation

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces/mocks"
)

func TestMatch_Validate_MatchesRegex_ShouldReturnNil(t *testing.T) {
	// setup
	validator := &Match{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "ID",
		},
		regex:            regexp.MustCompile(`^\d-\d\d\d$`),
		regexDescription: "ID matching #-###",
	}

	mMessageStore := &mocks.MessageStorer{}

	type testSubjectType struct {
		ID string
	}
	testSubject := &testSubjectType{ID: "1-123"}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.Nil(t, result)
	mMessageStore.AssertExpectations(t)
}

func TestMatch_Validate_DoesNotMatchRegex_ShouldReturnNil(t *testing.T) {
	// setup
	validator := &Match{
		AllFieldValidators: &AllFieldValidators{
			fieldName: "ID",
		},
		regex:            regexp.MustCompile(`^\d-\d\d\d$`),
		regexDescription: "ID matching #-###",
	}

	mMessageStore := &mocks.MessageStorer{}
	mMessageStore.On("GetMessage", mock.AnythingOfType("*validation.Match"), mock.AnythingOfType("*validation.ValidatorContext")).Return("Does Not Match")

	type testSubjectType struct {
		ID string
	}
	testSubject := &testSubjectType{ID: "asdf"}

	// test
	result := validator.Validate(testSubject, nil, mMessageStore)

	// assert
	assert.NotNil(t, result)
	mMessageStore.AssertExpectations(t)
}

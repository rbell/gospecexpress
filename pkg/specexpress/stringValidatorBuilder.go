package specexpress

import (
	"regexp"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// Matches provides a way to enforce a string contains a pattern defined by a regex
func (v *validatorBuilder) Matches(regex *regexp.Regexp, regexDescripton string) interfaces.ValidatorBuilder {
	vals := append(*v.validators, validation.NewMatch(v.fieldName, regex, regexDescripton))
	*v.validators = vals
	return v
}

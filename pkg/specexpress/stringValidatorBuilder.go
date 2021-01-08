// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specexpress

import (
	"regexp"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// Matches provides a way to enforce a string contains a pattern defined by a regex
func (v *validatorBuilder) Matches(regex *regexp.Regexp, regexDescripton string, options ...interfaces.ValidatorOption) interfaces.ValidatorBuilder {
	addValidator(v.validators, v.fieldName, v.fieldAlias, ApplyValidatorOptions(validation.NewMatch(v.fieldName, v.fieldAlias, regex, regexDescripton), options...))
	return v
}

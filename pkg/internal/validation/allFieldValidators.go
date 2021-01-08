// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import "gitlab.com/rbell/gospecexpress/pkg/interfaces"

// AllFieldValidators defines functionality shared across all Field Validators
type AllFieldValidators struct {
	fieldName            string
	fieldAlias           string
	overrideErrorMessage interfaces.MessageFormatter
}

// NewValidatorContext gets error message context for instance
func (a *AllFieldValidators) NewValidatorContext(instance interface{}, additionalContext map[string]interface{}) *ValidatorContext {
	m := map[string]interface{}{
		ContextFieldNameKey:  a.fieldName,
		ContextFieldAliasKey: a.fieldAlias,
		ContextInstanceKey:   instance,
	}
	if additionalContext != nil {
		m = mergeMap(m, additionalContext)
	}

	return NewValidatorMessageContext(instance, m)
}

// GetOverrideErrorMessage gets the overloaded message if overridden
func (a *AllFieldValidators) GetOverrideErrorMessage(ctx interfaces.ValidatorContextGetter) string {
	if a.overrideErrorMessage != nil {
		return a.overrideErrorMessage(ctx)
	}
	return ""
}

// SetOverrideErrorMessage gets the overloaded message if overridden
func (a *AllFieldValidators) SetOverrideErrorMessage(msgFormatter interfaces.MessageFormatter) {
	a.overrideErrorMessage = msgFormatter
}

func mergeMap(m1, m2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for k, v := range m1 {
		merged[k] = v
	}

	for k, v := range m2 {
		if k != ContextFieldNameKey && k != ContextFieldValueKey && k != ContextInstanceKey {
			merged[k] = v
		}
	}

	return merged
}

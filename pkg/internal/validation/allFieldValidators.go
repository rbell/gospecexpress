package validation

import "gitlab.com/rbell/gospecexpress/pkg/interfaces"

// AllFieldValidators defines functionality shared across all Field Validators
type AllFieldValidators struct {
	fieldName            string
	displayFieldName     string
	overrideErrorMessage func(ctx interfaces.ValidatorContextGetter) string
}

// NewValidatorContext gets error message context for instance
func (a *AllFieldValidators) NewValidatorContext(instance interface{}, additionalContext map[string]interface{}) *ValidatorContext {
	m := map[string]interface{}{
		ContextFieldNameKey:  a.fieldName,
		ContextFieldValueKey: a.displayFieldName,
	}
	if additionalContext != nil {
		m = mergeMap(m, additionalContext)
	}

	return NewValidatorMessageContext(instance, m)
}

func (a *AllFieldValidators) setOverrideErrorMessage(overLoad func(ctx interfaces.ValidatorContextGetter) string) {
	a.overrideErrorMessage = overLoad
}

// GetOverrideErrorMessage gets the overloaded message if overridden
func (a *AllFieldValidators) GetOverrideErrorMessage(ctx interfaces.ValidatorContextGetter) string {
	if a.overrideErrorMessage != nil {
		return a.overrideErrorMessage(ctx)
	}
	return ""
}

func mergeMap(m1, m2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for k, v := range m1 {
		merged[k] = v
	}

	for k, v := range m2 {
		if k != ContextFieldNameKey && k != ContextFieldValueKey {
			merged[k] = v
		}
	}

	return merged
}

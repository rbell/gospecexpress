package validation

// AllFieldValidators defines functionality shared across all Field Validators
type AllFieldValidators struct {
	FieldName        string
	DisplayFieldName string
}

// NewValidatorContext gets error message context for instance
func (a *AllFieldValidators) NewValidatorContext(instance interface{}, additionalContext map[string]interface{}) *ValidatorContext {
	m := map[string]interface{}{
		ContextFieldNameKey:  a.FieldName,
		ContextFieldValueKey: a.DisplayFieldName,
	}
	if additionalContext != nil {
		m = mergeMap(m, additionalContext)
	}

	return NewValidatorMessageContext(instance, m)
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

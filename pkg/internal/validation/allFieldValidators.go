package validation

// AllFieldValidators defines functionality shared across all Field Validators
type AllFieldValidators struct {
	FieldName        string
	DisplayFieldName string
}

// NewValidatorContext gets error message context for instance
func (a *AllFieldValidators) NewValidatorContext(instance interface{}, additionalContext ...interface{}) *ValidatorContext {
	ctx := []interface{}{a.FieldName, a.DisplayFieldName}
	ctx = append(ctx, additionalContext...)
	return NewErrorMessageContext(instance, ctx...)
}

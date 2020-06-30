package validators

import "gitlab.com/rbell/gospecexpress/pkg/errors"

// AllFieldValidators defines functionality shared across all Field Validators
type AllFieldValidators struct {
	FieldName        string
	DisplayFieldName string
}

// ErrorMessageContext gets error message context for instance
func (a *AllFieldValidators) ErrorMessageContext(instance interface{}, additionalContext ...interface{}) *errors.ErrorMessageContext {
	ctx := []interface{}{}
	ctx = append(ctx, a.FieldName, a.DisplayFieldName)
	ctx = append(ctx, additionalContext...)
	return errors.NewErrorMessageContext(instance, ctx...)
}

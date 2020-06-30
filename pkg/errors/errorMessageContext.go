package errors

import "gitlab.com/rbell/gospecexpress/internal/reflectionhelpers"

// ErrorMessageContext defines context for function that creates an error message
type ErrorMessageContext struct {
	Instance    interface{}
	ContextData []interface{}
}

// NewErrorMessageContext creates an initialized ErrorMessageContext
func NewErrorMessageContext(instance interface{}, context ...interface{}) *ErrorMessageContext {
	return &ErrorMessageContext{
		Instance:    instance,
		ContextData: context,
	}
}

// GetFieldValue gets the value of a specified field
func (e *ErrorMessageContext) GetFieldValue(fieldName string) interface{} {
	if v, ok := reflectionhelpers.GetFieldValue(e.Instance, fieldName); ok {
		return v.Interface()
	}
	return nil
}

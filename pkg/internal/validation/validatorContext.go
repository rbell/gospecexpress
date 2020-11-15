package validation

import "gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"

// ValidatorContext defines context for function that creates an error message
type ValidatorContext struct {
	instance    interface{}
	contextData []interface{}
}

// NewValidatorMessageContext creates an initialized ValidatorContext
func NewValidatorMessageContext(instance interface{}, context ...interface{}) *ValidatorContext {
	return &ValidatorContext{
		instance:    instance,
		contextData: context,
	}
}

// GetFieldValue gets the value of a specified field
func (e *ValidatorContext) GetFieldValue(fieldName string) interface{} {
	if v, ok := reflectionhelpers.GetFieldValue(e.instance, fieldName); ok {
		return v.Interface()
	}
	return nil
}

// GetContextData gets the context data set at time of validation of an instance
func (e *ValidatorContext) GetContextData() []interface{} {
	return e.contextData
}

// AddContextData adds data to the context
func (e *ValidatorContext) AddContextData(data interface{}) {
	e.contextData = append(e.contextData, data)
}

package validation

import (
	"sync"

	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"
)

const (
	// ContextFieldNameKey defines the context key name for the Field Name
	ContextFieldNameKey = "FieldName"
	// ContextFieldValueKey defines the context key name for the Field Value
	ContextFieldValueKey = "FieldValue"
)

// ValidatorContext defines context for function that creates an error message
type ValidatorContext struct {
	instance    interface{}
	contextData *sync.Map
}

// NewValidatorMessageContext creates an initialized ValidatorContext
func NewValidatorMessageContext(instance interface{}, data map[string]interface{}) *ValidatorContext {
	return &ValidatorContext{
		instance:    instance,
		contextData: mapToSyncMap(data),
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
func (e *ValidatorContext) GetContextData() map[string]interface{} {
	return syncMapToMap(e.contextData)
}

// AddContextData adds data to the context
func (e *ValidatorContext) AddContextData(key string, data interface{}) {
	e.contextData.Store(key, data)
}

func mapToSyncMap(m map[string]interface{}) *sync.Map {
	sm := &sync.Map{}
	for k, v := range m {
		sm.Store(k, v)
	}
	return sm
}

func syncMapToMap(sm *sync.Map) map[string]interface{} {
	m := make(map[string]interface{})
	sm.Range(func(key, value interface{}) bool {
		m[key.(string)] = value
		return true
	})
	return m
}

package catalog

const (
	ScopeContextKey = "ValidationScope"
)

// WithContext adds contextItems to the context
func WithContext(contextItems map[string]interface{}) func(something interface{}, context map[string]interface{}) {
	return func(something interface{}, context map[string]interface{}) {
		for k, item := range contextItems {
			context[k] = item
		}
	}
}

// WithContextItem adds contextItem with an explicitly defined key
func WithContextItem(key string, contextItem interface{}) func(something interface{}, context map[string]interface{}) {
	return func(something interface{}, context map[string]interface{}) {
		context[key] = contextItem
	}
}

// WithScope requests that the instance be validated for a specific scope
func WithScope(scope string) func(something interface{}, context map[string]interface{}) {
	return func(something interface{}, context map[string]interface{}) {
		context[ScopeContextKey] = scope
	}
}

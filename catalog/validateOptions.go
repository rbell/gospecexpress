package catalog

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

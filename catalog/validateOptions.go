package catalog

import "reflect"

// WithContext adds contextItems to the context, keyed by each item's type name.
// Perfect when passing a list of items where each item is of different type.
func WithContext(contextItems ...interface{}) func(something interface{}, context map[string]interface{}) {
	return func(something interface{}, context map[string]interface{}) {
		for _, item := range contextItems {
			key := reflect.TypeOf(item).Name()
			context[key] = item
		}
	}
}

// WithContextItem adds contextItem with an explicitly defined key
func WithContextItem(key string, contextItem interface{}) func(something interface{}, context map[string]interface{}) {
	return func(something interface{}, context map[string]interface{}) {
		context[key] = contextItem
	}
}

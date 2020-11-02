package interfaces

// ErrorMessageGetterFunc defines a function that, given error message context, return an error string
type ErrorMessageGetterFunc func(ctx ValidatorContextGetter) string

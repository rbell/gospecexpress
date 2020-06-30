package errors

// ErrorMessageGetterFunc defines a function that, given error message context, return an error string
type ErrorMessageGetterFunc func(ctx *ErrorMessageContext) string

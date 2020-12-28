package specexpress

import "gitlab.com/rbell/gospecexpress/pkg/interfaces"

// ApplyValidatorOptions applies options to a validator, returning the optioned validator
func ApplyValidatorOptions(v interfaces.Validator, options ...interfaces.ValidatorOption) interfaces.Validator {
	for _, o := range options {
		o(v)
	}
	return v
}

// OverrideMessageFromContext overrides the error message if the validator fails
func OverrideMessageFromContext(msgFormatter interfaces.MessageFormatter) func(validator interfaces.Validator) {
	return func(validator interfaces.Validator) {
		if overrider, ok := validator.(interfaces.MessageOverrider); ok {
			overrider.SetOverrideErrorMessage(msgFormatter)
		}
	}
}

// OverrideMessage overrides the error message if the validator fails
func OverrideMessage(msg string) func(validator interfaces.Validator) {
	return OverrideMessageFromContext(func(ctx interfaces.ValidatorContextGetter) string {
		return msg
	})
}

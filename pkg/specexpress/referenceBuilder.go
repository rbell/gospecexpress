package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// ValidateReference forces validating the value against the catalog (i.e. validate the address struct referenced by customer in the customer's Address field)
func (v *validatorBuilder) ValidateReference() interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.NewReferenceValidator(v.fieldName))
	*v.validators = vals
	return v
}

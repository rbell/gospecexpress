package builders

import (
	"gitlab.com/rbell/gospecexpress/internal/validators"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

// MaxLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MaxLength(length int) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validators.NewMaxLengthValidator(v.fieldName, length))
	*v.validators = vals
	return v
}

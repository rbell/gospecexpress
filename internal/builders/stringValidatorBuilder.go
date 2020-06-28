package builders

import (
	"gitlab.com/govalidate/internal/validators"
	"gitlab.com/govalidate/pkg/interfaces"
)

// MaxLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MaxLength(length int) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validators.NewMaxLengthValidator(v.fieldName, length))
	*v.validators = vals
	return v
}

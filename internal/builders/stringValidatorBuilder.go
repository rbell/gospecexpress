package builders

import (
	"gitlab.com/govalidate/internal/validators"
	"gitlab.com/govalidate/pkg/interfaces"
)

// MaxLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MaxLength(len int) interfaces.ValidatorBuilder {
	vals := append(*v.validators, validators.NewMaxLengthValidator(v.fieldName, len))
	*v.validators = vals
	return v
}

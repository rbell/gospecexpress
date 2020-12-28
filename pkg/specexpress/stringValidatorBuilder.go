package specexpress

import (
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/validation"
)

// MaxLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MaxLength(length int) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.NewMaxLengthValidator(v.fieldName, length))
	*v.validators = vals
	return v
}

// MinLength indicates a max length rule should be applied to field
func (v *validatorBuilder) MinLength(length int) interfaces.ValidatorBuilder {
	//nolint:gocritic // invalid
	vals := append(*v.validators, validation.NewMinLengthValidator(v.fieldName, length))
	*v.validators = vals
	return v
}

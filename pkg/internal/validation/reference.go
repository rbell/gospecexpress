package validation

import (
	"gitlab.com/rbell/gospecexpress/pkg/catalog"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
	"gitlab.com/rbell/gospecexpress/pkg/internal/reflectionhelpers"
)

// Reference validates a struct's reference to another type (i.e. Order.ShipAddress)
type Reference struct {
	*AllFieldValidators
	validationCatalog interfaces.Cataloger
}

// NewReferenceValidator returns an initialized Reference validator
func NewReferenceValidator(fieldName string) interfaces.Validator {
	return &Reference{
		AllFieldValidators: &AllFieldValidators{
			fieldName: fieldName,
		},
		validationCatalog: catalog.ValidationCatalog(),
	}
}

// Validate validates the reference
func (v *Reference) Validate(thing interface{}, contextData map[string]interface{}, _ interfaces.MessageStorer) error {
	// get the value of the reference and validate it
	if val, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		err := v.validationCatalog.ValidateWithContext(val.Interface(), contextData)
		if err != nil {
			return NewValidationError(v.fieldName, err.Error())
		}
	}

	return nil
}

package validators

import (
	"fmt"

	"gitlab.com/rbell/gospecexpress/pkg/specificationcatalog"

	"gitlab.com/rbell/gospecexpress/internal/reflectionhelpers"

	"gitlab.com/rbell/gospecexpress/pkg/errors"
	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

const defaultRequiredFieldMessage = "%v is required."

// RequiredField defines a validator requiring a field value be populated.
type RequiredField struct {
	fieldName string
}

// NewRequiredFieldValidator creates an initialized RequiredFieldValidator
func NewRequiredFieldValidator(fieldName string) interfaces.Validator {
	return &RequiredField{
		fieldName: fieldName,
	}
}

func init() {
	specificationcatalog.Catalog().MessageStore().SetMessage(&RequiredField{}, defaultRequiredFieldMessage)
}

// Validate validates the thing ensureing the field specified is populated
func (v *RequiredField) Validate(thing interface{}) error {
	if fv, ok := reflectionhelpers.GetFieldValue(thing, v.fieldName); ok {
		if fv.IsZero() {
			// TODO: Get message from a msg repository of some sorts
			// msg := catalog.Registry.GetMsg("en_US", "RequiredField", v.fieldName, forType)
			return errors.NewValidationError(v.fieldName, fmt.Sprintf(specificationcatalog.Catalog().MessageStore().GetMessage(v), v.fieldName))
		}
	}

	return nil
}

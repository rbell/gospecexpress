package reflectionhelpers

import (
	"reflect"
)

const (
	specTagKey = "spec"
)

// GetFieldAlias returns the tagged alias of the field.  If not tagged, fieldName is returned.
func GetFieldAlias(v reflect.Value, fieldName string) string {
	if sv, ok := StructValue(v); ok {
		t := sv.Type()
		if f, ok := t.FieldByName(fieldName); ok {
			if tag, ok := f.Tag.Lookup(specTagKey); ok {
				if tag != "" {
					return tag
				}
			}
		}
	}

	return fieldName
}

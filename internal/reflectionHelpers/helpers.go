package reflectionHelpers

import "reflect"

// StructValue gets the Structure value of a Value t.  If t is an interface wrapping a struct or pointer to a struct
// the struct referenced by the interface or pointer is returned.
func StructValue(t reflect.Value) (*reflect.Value, bool) {
	if t.Kind() == reflect.Struct {
		return &t, true
	}
	if t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		elem := t.Elem()
		if elem.Kind() == reflect.Struct {
			return &elem, true
		}
	}
	return nil, false
}

// GetFieldValue gets a field value from a struct if the field exists.
func GetFieldValue(thing interface{}, fieldName string) (*reflect.Value, bool) {
	tv := reflect.ValueOf(thing)
	if sv,ok := StructValue(tv); ok {
		if _,ok := sv.Type().FieldByName(fieldName); ok {
			fv := sv.FieldByName(fieldName)
			return &fv, true
		}
	}
	return nil, false
}

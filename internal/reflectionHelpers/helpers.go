package reflectionhelpers

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errBadComparisonType = errors.New("invalid type for comparison")
	errBadComparison     = errors.New("incompatible types for comparison")
	errNoComparison      = errors.New("missing argument for comparison")
)

var zero reflect.Value

type kind int

const (
	invalidKind kind = iota
	boolKind
	complexKind
	intKind
	floatKind
	stringKind
	uintKind
)

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
	if sv, ok := StructValue(tv); ok {
		if _, ok := sv.Type().FieldByName(fieldName); ok {
			fv := sv.FieldByName(fieldName)
			return &fv, true
		}
	}
	return nil, false
}

// Eq evaluates the comparison a == b || a == c || ...
func Eq(arg1 reflect.Value, arg2 ...reflect.Value) (bool, error) {
	v1 := indirectInterface(arg1)
	if v1 != zero {
		if t1 := v1.Type(); !t1.Comparable() {
			return false, fmt.Errorf("uncomparable type %s: %v", t1, v1)
		}
	}
	if len(arg2) == 0 {
		return false, errNoComparison
	}
	//nolint:errcheck // not interested in err
	k1, _ := basicKind(v1)
	for _, arg := range arg2 {
		//nolint:errcheck // not interested in err
		v2 := indirectInterface(arg)
		//nolint:errcheck // not interested in err
		k2, _ := basicKind(v2)
		truth := false
		if k1 != k2 {
			// Special case: Can compare integer values regardless of type's sign.
			switch {
			case k1 == intKind && k2 == uintKind:
				truth = v1.Int() >= 0 && uint64(v1.Int()) == v2.Uint()
			case k1 == uintKind && k2 == intKind:
				truth = v2.Int() >= 0 && v1.Uint() == uint64(v2.Int())
			default:
				return false, errBadComparison
			}
		} else {
			switch k1 {
			case boolKind:
				truth = v1.Bool() == v2.Bool()
			case complexKind:
				truth = v1.Complex() == v2.Complex()
			case floatKind:
				truth = v1.Float() == v2.Float()
			case intKind:
				truth = v1.Int() == v2.Int()
			case stringKind:
				truth = v1.String() == v2.String()
			case uintKind:
				truth = v1.Uint() == v2.Uint()
			default:
				if v2 == zero {
					truth = v1 == v2
				} else {
					if t2 := v2.Type(); !t2.Comparable() {
						return false, fmt.Errorf("uncomparable type %s: %v", t2, v2)
					}
					truth = v1.Interface() == v2.Interface()
				}
			}
		}
		if truth {
			return true, nil
		}
	}
	return false, nil
}

// Ne evaluates the comparison a != b.
func Ne(arg1, arg2 reflect.Value) (bool, error) {
	// != is the inverse of ==.
	equal, err := Eq(arg1, arg2)
	return !equal, err
}

// Lt evaluates the comparison a < b.
func Lt(arg1, arg2 reflect.Value) (bool, error) {
	v1 := indirectInterface(arg1)
	k1, err := basicKind(v1)
	if err != nil {
		return false, err
	}
	v2 := indirectInterface(arg2)
	k2, err := basicKind(v2)
	if err != nil {
		return false, err
	}
	truth := false
	if k1 != k2 {
		// Special case: Can compare integer values regardless of type's sign.
		switch {
		case k1 == intKind && k2 == uintKind:
			truth = v1.Int() < 0 || uint64(v1.Int()) < v2.Uint()
		case k1 == uintKind && k2 == intKind:
			truth = v2.Int() >= 0 && v1.Uint() < uint64(v2.Int())
		default:
			return false, errBadComparison
		}
	} else {
		switch k1 {
		case boolKind, complexKind:
			return false, errBadComparisonType
		case floatKind:
			truth = v1.Float() < v2.Float()
		case intKind:
			truth = v1.Int() < v2.Int()
		case stringKind:
			truth = v1.String() < v2.String()
		case uintKind:
			truth = v1.Uint() < v2.Uint()
		default:
			panic("invalid kind")
		}
	}
	return truth, nil
}

// Le evaluates the comparison <= b.
func Le(arg1, arg2 reflect.Value) (bool, error) {
	// <= is < or ==.
	lessThan, err := Lt(arg1, arg2)
	if lessThan || err != nil {
		return lessThan, err
	}
	return Eq(arg1, arg2)
}

// Gt evaluates the comparison a > b.
func Gt(arg1, arg2 reflect.Value) (bool, error) {
	// > is the inverse of <=.
	lessOrEqual, err := Le(arg1, arg2)
	if err != nil {
		return false, err
	}
	return !lessOrEqual, nil
}

// Ge evaluates the comparison a >= b.
func Ge(arg1, arg2 reflect.Value) (bool, error) {
	// >= is the inverse of <.
	lessThan, err := Lt(arg1, arg2)
	if err != nil {
		return false, err
	}
	return !lessThan, nil
}

// indirectInterface returns the concrete value in an interface value,
// or else the zero reflect.Value.
// That is, if v represents the interface value x, the result is the same as reflect.ValueOf(x):
// the fact that x was an interface value is forgotten.
func indirectInterface(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Interface {
		return v
	}
	if v.IsNil() {
		return reflect.Value{}
	}
	return v.Elem()
}

func basicKind(v reflect.Value) (kind, error) {
	switch v.Kind() {
	case reflect.Bool:
		return boolKind, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intKind, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintKind, nil
	case reflect.Float32, reflect.Float64:
		return floatKind, nil
	case reflect.Complex64, reflect.Complex128:
		return complexKind, nil
	case reflect.String:
		return stringKind, nil
	}
	return invalidKind, errBadComparisonType
}

package gospecexpress

import (
	"sync"

	"github.com/rbell/gospecexpress/interfaces"
)

type fieldValidator struct {
	isOptional bool
	alias      string
	validators []interfaces.Validator
	condition  interfaces.ValidationPredicate
	mux        *sync.Mutex
}

func (v *fieldValidator) addValidator(validator interfaces.Validator) {
	v.mux.Lock()
	defer v.mux.Unlock()
	v.validators = append(v.validators, validator)
}

func setOptional(fieldValidators *sync.Map, fieldName string, optional bool) {
	var fv *fieldValidator
	if v, ok := fieldValidators.Load(fieldName); ok {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fv, _ = v.(*fieldValidator)
	} else {
		fv = &fieldValidator{
			isOptional: false,
			validators: []interfaces.Validator{},
			mux:        &sync.Mutex{},
		}
	}
	fv.isOptional = optional
	fieldValidators.Store(fieldName, fv)
}

func setCondition(fieldValidators *sync.Map, fieldName string, condition interfaces.ValidationPredicate) {
	var fv *fieldValidator
	if v, ok := fieldValidators.Load(fieldName); ok {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fv, _ = v.(*fieldValidator)
	} else {
		fv = &fieldValidator{
			isOptional: false,
			validators: []interfaces.Validator{},
			mux:        &sync.Mutex{},
		}
	}
	fv.condition = condition
	fieldValidators.Store(fieldName, fv)
}

package gospecexpress

import (
	"sync"

	"github.com/rbell/gospecexpress/interfaces"
)

type fieldExpression struct {
	isOptional bool
	alias      string
	validators []interfaces.Validator
	condition  interfaces.FieldValidationCondition
	mux        *sync.Mutex
}

func (v *fieldExpression) addValidator(validator interfaces.Validator) {
	v.mux.Lock()
	defer v.mux.Unlock()
	v.validators = append(v.validators, validator)
}

func setCondition(fieldExpressions *sync.Map, fieldName string, condition interfaces.FieldValidationCondition) {
	var fe *fieldExpression
	if v, ok := fieldExpressions.Load(fieldName); ok {
		//nolint:errcheck // We are in control of key and value types so should no need to check error
		fexps, _ := v.([]*fieldExpression)
		fe = fexps[len(fexps)-1]
	} else {
		fe = &fieldExpression{
			isOptional: false,
			validators: []interfaces.Validator{},
			mux:        &sync.Mutex{},
		}
		fieldExpressions.Store(fieldName, []*fieldExpression{fe})
	}
	fe.condition = condition
}

package errormessagestore

import (
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/errors"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

var _ interfaces.MessageStorer = &defaultMessageStore{}

type defaultMessageStore struct {
	messages map[string]errors.ErrorMessageGetterFunc
}

// NewDefaultMessageStore returns an initialized defaultMessageStore
func NewDefaultMessageStore() interfaces.MessageStorer {
	return &defaultMessageStore{
		messages: make(map[string]errors.ErrorMessageGetterFunc),
	}
}

// GetMessage gets a message for a validator
func (d *defaultMessageStore) GetMessage(validator interfaces.Validator, ctx *errors.ErrorMessageContext) string {
	if m, ok := d.messages[typeKey(validator)]; ok {
		return m(ctx)
	}
	return ""
}

// SetMessage stores a message for a validator to use
func (d *defaultMessageStore) SetMessage(validator interfaces.Validator, getterFunc errors.ErrorMessageGetterFunc) {
	d.messages[typeKey(validator)] = getterFunc
}

func typeKey(validator interfaces.Validator) string {
	t := reflect.TypeOf(validator).Elem()
	return t.PkgPath() + "|" + t.Name()
}

package errormessagestore

import (
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

var _ interfaces.MessageStorer = &defaultMessageStore{}

type defaultMessageStore struct {
	messages map[string]string
}

// NewDefaultMessageStore returns an initialized defaultMessageStore
func NewDefaultMessageStore() interfaces.MessageStorer {
	return &defaultMessageStore{
		messages: make(map[string]string),
	}
}

// GetMessage gets a message for a validator
func (d *defaultMessageStore) GetMessage(validator interfaces.Validator) string {
	if m, ok := d.messages[typeKey(validator)]; ok {
		return m
	}
	return ""
}

// StoreMessage stores a message for a validator to use
func (d *defaultMessageStore) StoreMessage(validator interfaces.Validator, msg string) {
	d.messages[typeKey(validator)] = msg
}

func typeKey(validator interfaces.Validator) string {
	t := reflect.TypeOf(validator).Elem()
	return t.PkgPath() + "|" + t.Name()
}

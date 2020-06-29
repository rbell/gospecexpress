package errormessagestore

import "gitlab.com/rbell/gospecexpress/pkg/interfaces"

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
func (d *defaultMessageStore) GetMessage(forValidator string) string {
	if m, ok := d.messages[forValidator]; ok {
		return m
	}
	return ""
}

// StoreMessage stores a message for a validator to use
func (d *defaultMessageStore) StoreMessage(forValidator, msg string) {
	d.messages[forValidator] = msg
}

// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package catalog

import (
	"reflect"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

var _ interfaces.MessageStorer = &defaultMessageStore{}

type defaultMessageStore struct {
	messages map[string]interfaces.ErrorMessageGetterFunc
}

// NewDefaultMessageStore returns an initialized defaultMessageStore
func NewDefaultMessageStore() interfaces.MessageStorer {
	return &defaultMessageStore{
		messages: make(map[string]interfaces.ErrorMessageGetterFunc),
	}
}

// GetMessage gets a message for a validator
func (d *defaultMessageStore) GetMessage(validator interfaces.Validator, ctx interfaces.ValidatorContextGetter) string {
	if overrider, ok := validator.(interfaces.MessageOverrider); ok {
		if msg := overrider.GetOverrideErrorMessage(ctx); msg != "" {
			return msg
		}
	}
	if m, ok := d.messages[typeKey(validator)]; ok {
		return m(ctx)
	}
	return ""
}

// SetMessage stores a message for a validator to use
func (d *defaultMessageStore) SetMessage(validator interfaces.Validator, getterFunc interfaces.ErrorMessageGetterFunc) {
	d.messages[typeKey(validator)] = getterFunc
}

func typeKey(validator interfaces.Validator) string {
	t := reflect.TypeOf(validator).Elem()
	return t.PkgPath() + "|" + t.Name()
}

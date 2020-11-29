package xredis

import (
	"encoding/json"
)

// String xredis string
type String string

// Unmarshal ..
func (t String) Unmarshal(obj interface{}) error {
	return json.Unmarshal([]byte(t), obj)
}

// String ..
func (t String) String() string {
	return string(t)
}

type CallbackHandler func() (interface{}, error)

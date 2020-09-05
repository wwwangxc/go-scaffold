package util

import uuid "github.com/satori/go.uuid"

// GenUUID ..
func GenUUID() string {
	return uuid.NewV4().String()
}

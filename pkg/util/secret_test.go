package util

import (
	"testing"
)

func TestSecret(t *testing.T) {
	str := "123456"
	t.Log(WithSecret(str, "盐"))
}

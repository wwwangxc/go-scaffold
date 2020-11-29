package config

import "testing"

func TestConfig(t *testing.T) {
	Init()
	t.Log(AllKeys())
}

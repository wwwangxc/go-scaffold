package log

import "testing"

func TestInfo(t *testing.T) {
	Info("infomation")
	Panic("panic infomation")
}

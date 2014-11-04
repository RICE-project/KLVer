package logger

import "testing"

var a Logger

func TestLogger(t *testing.T) {
	err := a.SetNewLogger()
	if err != nil {
		t.Error("Err: ", err)
	}
	defer a.CloseLogger()
	a.SetPrefix("TEST")
	a.logTest("The quick brown fox jumps over the lazy dog.")
	a.logTest("RICE.lib.logger.Logger test.", a)
}

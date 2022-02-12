package solution

import (
	"testing"

	"github.com/kyokomi/emoji"
)

func TestGetMessage(t *testing.T) {
	if msg := GetMessage(); msg != emoji.Sprint(message) {
		t.Errorf("Wrong answer: msg: %s", msg)
	}
}

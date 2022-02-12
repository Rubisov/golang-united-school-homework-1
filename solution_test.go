package solution

import (
	"github.com/kyokomi/emoji/v2"
	"testing"
)

func TestGetMessage(t *testing.T) {
	result := GetMessage()
	GetMessage := emoji.Sprint("Hello :world_map:!")
	if result != GetMessage {
		t.Errorf("Result string is incorrect: got %s, want %s", result, GetMessage)
	}
}

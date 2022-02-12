package solution

import (
	"github.com/kyokomi/emoji/v2"
	"testing"
)

func WorldMessage(t *testing.T) {
	result := helloWorld()
	WorldMessage := emoji.Sprint("Hello :world_map:!")
	if result != WorldMessage {
		t.Errorf("Result string is incorrect: got %s, want %s", result, WorldMessage)
	}
}

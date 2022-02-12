package main

import (
	"testing"
	"github.com/kyokomi/emoji/v2"
)

func testSum(t *testing.T) {
	result := helloWorld()
	worldMessage := emoji.Sprint("Hello :world_map:!")
	if result != worldMessage {
		t.Errorf("Result string is incorrect: got %s, want %s", result, worldMessage)
	}
}
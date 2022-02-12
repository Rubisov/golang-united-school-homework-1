package main

import (
	"github.com/kyokomi/emoji/v2"
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	GetMessage := emoji.Sprint("Hello :world_map:!")
	expected := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})
	msg := GetMessage()

	if !strings.EqualFold(msg, expected) {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", expected, msg)
	}
}

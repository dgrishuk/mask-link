package main

import (
	"testing"
)

func TestMaskLinks(t *testing.T) {
	testMessage := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	expected := "Here's my spammy page: http://******************* see you."

	result := MaskLink(testMessage)

	if result != expected {
		t.Errorf("Incorrect result. Expect %s, got %s\n", expected, result)
	}

}

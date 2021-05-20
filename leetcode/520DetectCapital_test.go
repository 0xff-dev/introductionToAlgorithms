package leetcode

import "testing"

func TestDetectCapitalUse(t *testing.T) {
	word := "USA"
	if !detectCapitalUse(word) {
		t.Fatalf("expect true get false")
	}

	word = "FlaG"
	if detectCapitalUse(word) {
		t.Fatalf("expect false get true")
	}
}

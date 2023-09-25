package leetcode

import "testing"

func TestCharacterReplacement(t *testing.T) {
	s := "ABAB"
	k := 2
	if r := characterReplacement(s, k); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	s = "AABABBA"
	k = 1
	if r := characterReplacement(s, k); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

package leetcode

import "testing"

func TestNumberofSubstrings(t *testing.T) {
	s := "abcabc"
	if r := numberOfSubstrings(s); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	s = "aaacb"
	if r := numberOfSubstrings(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	s = "abc"
	if r := numberOfSubstrings(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

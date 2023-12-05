package leetcode

import "testing"

func TestNumberOfMatches(t *testing.T) {
	if r := numberOfMatches(7); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	if r := numberOfMatches(14); r != 13 {
		t.Fatalf("expect 6 get %d", r)
	}
}

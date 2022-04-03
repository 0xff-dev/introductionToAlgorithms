package leetcode

import "testing"

func TestMinWindow(t *testing.T) {
	s, t1 := "ADOBECODEBANC", "ABC"
	if r := minWindow(s, t1); r != "BANC" {
		t.Fatalf("expect 'BANC' get %s", r)
	}
}

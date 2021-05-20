package leetcode

import "testing"

func TestFindLUSLength(t *testing.T) {
	a, b := "aba", "cdc"
	if r := findLUSlength(a, b); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	a, b = "aaa", "bbb"
	if r := findLUSlength(a, b); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

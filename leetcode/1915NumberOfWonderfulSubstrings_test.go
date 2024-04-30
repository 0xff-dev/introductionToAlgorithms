package leetcode

import "testing"

func TestWonderfulSubstrings(t *testing.T) {
	word := "aba"
	exp := int64(4)
	if r := wonderfulSubstrings(word); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

package leetcode

import "testing"

func TestTakeCharacters(t *testing.T) {
	s, k, exp := "aabaaaacaabc", 2, 8
	if r := takeCharacters(s, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

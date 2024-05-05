package leetcode

import "testing"

func TestMinimumDeletions(t *testing.T) {
	word := "aabcaba"
	k := 0
	exp := 3
	if r := minimumDeletions(word, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	word = "dabdcbdcdcd"
	k = 2
	exp = 2
	if r := minimumDeletions(word, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	word = "aaabaaa"
	k = 2
	exp = 1
	if r := minimumDeletions(word, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

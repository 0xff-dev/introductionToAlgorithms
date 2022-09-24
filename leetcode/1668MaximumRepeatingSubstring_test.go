package leetcode

import "testing"

func TestMaxRepeating(t *testing.T) {
	seq, word := "ababc", "ab"
	if r := maxRepeating(seq, word); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	seq, word = "ababc", "ba"
	if r := maxRepeating(seq, word); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	seq, word = "ababc", "ac"
	if r := maxRepeating(seq, word); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	seq, word = "aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"

}

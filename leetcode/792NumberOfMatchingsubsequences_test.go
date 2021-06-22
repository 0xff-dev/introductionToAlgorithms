package leetcode

import "testing"

func TestNumMatchingSubseq(t *testing.T) {
	s, words := "abcde", []string{"a", "bb", "acd", "ace"}
	if r := numMatchingSubseq(s, words); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	s, words = "dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}
	if r := numMatchingSubseq(s, words); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

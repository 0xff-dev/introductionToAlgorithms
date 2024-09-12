package leetcode

import "testing"

func TestCountConsistentStrings(t *testing.T) {
	allowed, words, exp := "ab", []string{"ad", "bd", "aaab", "baa", "badab"}, 2
	if r := countConsistentStrings(allowed, words); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	allowed, words, exp = "abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, 7
	if r := countConsistentStrings(allowed, words); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	allowed, words, exp = "cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, 4
	if r := countConsistentStrings(allowed, words); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

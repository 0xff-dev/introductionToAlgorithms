package leetcode

import "testing"

func TestMaxProduct318(t *testing.T) {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	if r := maxProduct318(words); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}

	words = []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	if r := maxProduct318(words); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	words = []string{"a", "aa", "aaa", "aaaa"}
	if r := maxProduct318(words); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

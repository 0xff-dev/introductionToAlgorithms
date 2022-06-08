package leetcode

import "testing"

func TestRemovePalindromicSub(t *testing.T) {
	s := "abab"
	if r := removePalindromeSub(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s = "ababbaba"
	if r := removePalindromeSub(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	s = "abb"
	if r := removePalindromeSub(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s = "baabb"
	if r := removePalindromeSub(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

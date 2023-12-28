package leetcode

import "testing"

func TestGetLengthOfOptimalCompression(t *testing.T) {
	s := "aaabcccd"
	k := 2
	if r := getLengthOfOptimalCompression(s, k); r != 4 {
		t.Fatalf("input: %s, k: %d expect 4 get %d", s, k, r)
	}
	s = "aabbaa"
	k = 2
	if r := getLengthOfOptimalCompression(s, k); r != 2 {
		t.Fatalf("input: %s, k: %d expect 2 get %d", s, k, r)
	}
	s = "aaaaaaaaaaa"
	k = 0
	if r := getLengthOfOptimalCompression(s, k); r != 3 {
		t.Fatalf("input: %s, k: %d, expect 3 get %d", s, k, r)
	}
}

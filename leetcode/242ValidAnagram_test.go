package leetcode

import "testing"

func TestIsAnagram(t *testing.T) {
	s, t1 := "anagram", "nagaram"
	if !isAnagram(s, t1) {
		t.Fatalf("expect true get false")
	}
}

package leetcode

import "testing"

func TestLongestSubstring(t *testing.T) {
	s := "aaabb"
	k := 3
	if r := longestSubstring(s, k); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	s = "ababbc"
	k = 2
	if r := longestSubstring(s, k); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	s = "bbaaacbd"
	k = 3
	if r := longestSubstring(s, k); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	s = "aaabbjerwjgjrjfdhdjkjgiorjvkofjiotjrijfifjijdifjoidhjtirutiohfgjdogjhierhtnjgre"
	k = 2
	if r := longestSubstring(s, k); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}

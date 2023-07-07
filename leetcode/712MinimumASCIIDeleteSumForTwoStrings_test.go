package leetcode

import "testing"

func TestMinimumDeleteSum(t *testing.T) {
	s1, s2 := "sea", "eat"
	if r := minimumDeleteSum(s1, s2); r != 231 {
		t.Fatalf("expect 231 get %d", r)
	}
	s1, s2 = "delete", "leet"
	if r := minimumDeleteSum(s1, s2); r != 403 {
		t.Fatalf("expect 403 get %d", r)
	}
}

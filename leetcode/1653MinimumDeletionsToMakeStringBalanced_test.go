package leetcode

import "testing"

func TestMinimumDeletions1653(t *testing.T) {
	s, exp := "aababbab", 2
	if r := minimumDeletions1653(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = "bbaaaaabb", 2
	if r := minimumDeletions1653(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

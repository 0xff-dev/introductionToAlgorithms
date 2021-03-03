package leetcode

import "testing"

func TestClimbStairs(t *testing.T) {
	n := 3
	if r := climbStairs(n); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

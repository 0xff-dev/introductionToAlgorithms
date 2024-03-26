package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{1, 2, 0}
	exp := 3
	if r := firstMissingPositive(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

package leetcode

import "testing"

func TestArraySign(t *testing.T) {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	if r := arraySign(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

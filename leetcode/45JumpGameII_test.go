package leetcode

import "testing"

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	if r := jump(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

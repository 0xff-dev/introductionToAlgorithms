package leetcode

import "testing"

func TestTriangleNumber(t *testing.T) {
	nums := []int{2, 2, 3, 4}
	if r := triangleNumber(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	nums = []int{4, 2, 3, 4}
	if r := triangleNumber(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

package leetcode

import "testing"

func TestMaxRotateFunction(t *testing.T) {
	nums := []int{4, 3, 2, 6}
	if r := maxRotateFunction(nums); r != 26 {
		t.Fatalf("expect 26 get %d", r)
	}

	nums = []int{100}
	if r := maxRotateFunction(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

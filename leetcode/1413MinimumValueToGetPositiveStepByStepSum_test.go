package leetcode

import "testing"

func TestMinStartValue(t *testing.T) {
	nums := []int{-3, 2, -3, 4, 2}
	if r := minStartValue(nums); r != 5 {
		t.Fatalf("expect %d get %d", 5, r)
	}

	nums = []int{1, 2}
	if r := minStartValue(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

package leetcode

import "testing"

func TestNumberOfArithmeticSlices(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	if r := numberOfArithmeticSlices413(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	nums = []int{1}
	if r := numberOfArithmeticSlices413(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

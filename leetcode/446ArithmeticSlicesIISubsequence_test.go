package leetcode

import (
	"testing"
)

func TestNumberOfArithmeticsSlices(t *testing.T) {
	nums := []int{2, 4, 6, 8, 10}
	if r := numberOfArithmeticSlices(nums); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	nums = []int{7, 7, 7, 7, 7}
	if r := numberOfArithmeticSlices(nums); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}
}

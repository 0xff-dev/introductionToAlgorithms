package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 3, 2}
	if r := singleNumber(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

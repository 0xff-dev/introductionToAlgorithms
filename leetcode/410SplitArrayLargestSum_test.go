package leetcode

import "testing"

func TestSplitArray(t *testing.T) {
	nums := []int{7, 2, 5, 10, 8}
	k := 2
	exp := 18
	if r := splitArray(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	nums = []int{1, 2, 3, 4, 5}
	k = 2
	exp = 9
	if r := splitArray(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{2, 3, 1, 2, 4, 3}
	k = 5
	exp = 4
	if r := splitArray(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

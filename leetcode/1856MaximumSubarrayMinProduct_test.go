package leetcode

import "testing"

func TestMaxSumMinProduct(t *testing.T) {
	nums := []int{2, 3, 3, 1, 2}
	if r := maxSumMinProduct(nums); r != 18 {
		t.Fatalf("expect 18 get %d", r)
	}

	nums = []int{3, 1, 5, 6, 4, 2}
	if r := maxSumMinProduct(nums); r != 60 {
		t.Fatalf("expect 60 get %d", r)
	}

	nums = []int{2, 5, 4, 2, 4, 5, 3, 1, 2}
	if r := maxSumMinProduct(nums); r != 50 {
		t.Fatalf("expect 50 get %d", r)
	}
}

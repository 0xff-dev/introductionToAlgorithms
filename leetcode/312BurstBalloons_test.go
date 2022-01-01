package leetcode

import (
	"testing"
)

func TestMaxConins(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	if r := maxCoins(nums); r != 167 {
		t.Fatalf("expect 167 get %d", r)
	}

	nums = []int{1, 5}
	if r := maxCoins(nums); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}

	nums = []int{2, 1, 3, 2, 4}
	if r := maxCoins(nums); r != 66 {
		t.Fatalf("expect 66 get %d", r)
	}

	nums = []int{2, 1, 3, 2, 2, 4}
	if r := maxCoins(nums); r != 82 {
		t.Fatalf("expect 82 get %d", r)
	}
}

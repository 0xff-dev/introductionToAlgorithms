package leetcode

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	if r := findMin(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	nums = []int{3, 4, 5, 1, 2}
	if r := findMin(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	nums = []int{4, 5, 6, 7, 0}
	if r := findMin(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	nums = []int{11, 13, 15, 17}
	if r := findMin(nums); r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}

}

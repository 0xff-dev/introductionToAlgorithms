package leetcode

import "testing"

func TestFindNumberOfLIS(t *testing.T) {
	nums := []int{1, 3, 5, 4, 7}
	if r := findNumberOfLIS(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	nums = []int{2, 2, 2, 2, 2}
	if r := findNumberOfLIS(nums); r != 5 {
		t.Fatalf("epxect 5 get %d", r)
	}
	nums = []int{1, 2, 3}
	if r := findNumberOfLIS(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	nums = []int{1, 3, 5, 4, 7, 7, 8, 7, 2, 10, 19, 23, 10, 101, 234, 78, 23}
	if r := findNumberOfLIS(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	nums = []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	if r := findNumberOfLIS(nums); r != 27 {
		t.Fatalf("expect 27 get %d", r)
	}
}

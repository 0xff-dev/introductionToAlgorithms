package leetcode

import "testing"

func TestMaximumValueSum(t *testing.T) {
	nums := []int{1, 2, 1}
	k := 3
	edges := [][]int{{0, 1}, {0, 2}}
	exp := int64(6)
	if r := maximumValueSum(nums, k, edges); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{2, 3}
	k = 7
	edges = [][]int{{0, 1}}
	exp = 9
	if r := maximumValueSum(nums, k, edges); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{7, 7, 7, 7, 7, 7}
	k = 3
	edges = [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}
	exp = 42
	if r := maximumValueSum(nums, k, edges); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

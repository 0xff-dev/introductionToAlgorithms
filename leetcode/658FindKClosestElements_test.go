package leetcode

import "testing"

func TestFindClosestElements(t *testing.T) {
	arr, k, x := []int{1, 2, 3, 4, 5}, 4, 3
	r := findClosestElements(arr, k, x)
	t.Logf("%v", r)
	arr, k, x = []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5
	t.Logf("%v", findClosestElements(arr, k, x))
	arr, k, x = []int{3, 5, 8, 10}, 2, 15
	t.Logf("%v", s2(arr, k, x))
}

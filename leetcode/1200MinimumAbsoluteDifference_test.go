package leetcode

import "testing"

func TestMinimumAbsDifference(t *testing.T) {
	arr := []int{4, 2, 1, 3}
	r := minimumAbsDifference(arr)
	t.Logf("%v ", r)
	arr = []int{1, 3, 6, 10, 15}
	r = minimumAbsDifference(arr)
	t.Logf("%v ", r)
	arr = []int{3, 8, -10, 23, 19, -4, -14, 27}
	r = minimumAbsDifference(arr)
	t.Logf("%v ", r)
}

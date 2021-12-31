package leetcode

import "testing"

func TestConstruct2DArray(t *testing.T) {
	original := []int{1, 2, 3, 4}
	r := construct2DArray(original, 2, 2)
	t.Logf("%v ", r)

	original = []int{1, 2}
	r = construct2DArray(original, 1, 1)
	t.Logf("%v ", r)

	original = []int{1, 2, 3}
	r = construct2DArray(original, 1, 3)
	t.Logf("%v ", r)
}

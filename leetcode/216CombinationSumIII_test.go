package leetcode

import "testing"

func TestCombinationSum3(t *testing.T) {
	k, n := 3, 7
	r := combinationSum3(k, n)
	t.Logf("%v", r)

	k, n = 3, 9
	r = combinationSum3(k, n)
	t.Logf("%v", r)
	k, n = 4, 1
	r = combinationSum3(k, n)
	t.Logf("%v", r)
}

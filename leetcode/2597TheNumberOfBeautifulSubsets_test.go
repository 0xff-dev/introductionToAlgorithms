package leetcode

import "testing"

func TestBeautifulSubsets(t *testing.T) {
	nums := []int{
		4, 2, 5, 9, 10, 3,
	}
	k := 1
	exp := 23
	if r := beautifulSubsets(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

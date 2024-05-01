package leetcode

import "testing"

func TestLargestSumofAverages(t *testing.T) {
	nums := []int{9, 1, 2, 3, 9}
	k := 3
	exp := 20.000
	if r := largestSumOfAverages(nums, k); r != exp {
		t.Fatalf("expect %.5f get %.5f", exp, r)
	}
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 4
	exp = 20.50000
	if r := largestSumOfAverages(nums, k); r != exp {
		t.Fatalf("expect %.5f get %.5f", exp, r)
	}
}

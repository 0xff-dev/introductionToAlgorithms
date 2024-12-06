package leetcode

import "testing"

func TestMaxCount(t *testing.T) {
	banned, n, maxSum, exp := []int{1, 6, 5}, 5, 6, 2
	if r := maxCount(banned, n, maxSum); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	banned, n, maxSum, exp = []int{1, 2, 3, 4, 5, 6, 7}, 8, 1, 0
	if r := maxCount(banned, n, maxSum); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	banned, n, maxSum, exp = []int{11}, 7, 50, 7
	if r := maxCount(banned, n, maxSum); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

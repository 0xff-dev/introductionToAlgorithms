package leetcode

import "testing"

func TestMaxSumAfterPartitioning(t *testing.T) {
	arr := []int{1, 15, 7, 9, 2, 5, 10}
	k := 3
	exp := 84
	if r := maxSumAfterPartitioning(arr, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	arr = []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}
	k = 4
	exp = 83
	if r := maxSumAfterPartitioning(arr, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	arr = []int{1}
	k = 1
	exp = 1
	if r := maxSumAfterPartitioning(arr, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

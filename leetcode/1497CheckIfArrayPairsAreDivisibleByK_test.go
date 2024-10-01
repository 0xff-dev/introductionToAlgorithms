package leetcode

import "testing"

func TestCanArrange(t *testing.T) {
	arr, k, exp := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5, true
	if r := canArrange(arr, k); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	arr, k, exp = []int{1, 2, 3, 4, 5, 6}, 7, true
	if r := canArrange(arr, k); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
	arr, k, exp = []int{1, 2, 3, 4, 5, 6}, 10, false
	if r := canArrange(arr, k); r != exp {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

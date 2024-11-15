package leetcode

import "testing"

func TestFindLenghtOfShortestSubarray(t *testing.T) {
	arr, exp := []int{1, 2, 3, 10, 4, 2, 3, 5}, 3
	if r := findLengthOfShortestSubarray(arr); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	arr, exp = []int{5, 4, 3, 2, 1}, 4
	if r := findLengthOfShortestSubarray(arr); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	arr, exp = []int{2, 2, 2, 1, 1, 1}, 3
	if r := findLengthOfShortestSubarray(arr); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

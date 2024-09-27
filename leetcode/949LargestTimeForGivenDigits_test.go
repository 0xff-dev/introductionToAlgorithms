package leetcode

import "testing"

func TestLargestTimeFromDigits(t *testing.T) {
	arr, exp := []int{1, 2, 3, 4}, "23:41"
	if r := largestTimeFromDigits(arr); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	arr, exp = []int{5, 5, 5, 5}, ""
	if r := largestTimeFromDigits(arr); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}

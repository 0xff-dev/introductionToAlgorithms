package leetcode

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	exp := 5
	if r := findUnsortedSubarray(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{1, 2, 3, 4}
	exp = 0
	if r := findUnsortedSubarray(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{1}
	if r := findUnsortedSubarray(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

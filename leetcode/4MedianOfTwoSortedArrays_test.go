package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	n1 := []int{1, 3}
	n2 := []int{2}
	if r := findMedianSortedArrays(n1, n2); r != 2 {
		t.Fatalf("expect 2 get %f", r)
	}
	n1 = []int{1, 2}
	n2 = []int{3, 4}
	if r := findMedianSortedArrays(n1, n2); r != 2.5 {
		t.Fatalf("expect 2.5 get %f", r)
	}
}

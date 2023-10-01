package leetcode

import "testing"

func TestFourSumCount(t *testing.T) {
	n1, n2, n3, n4 := []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}
	if r := fourSumCount(n1, n2, n3, n4); r != 2 {
		t.Fatalf("expect %d get %d", 2, r)
	}
	n1, n2, n3, n4 = []int{0}, []int{0}, []int{0}, []int{0}
	if r := fourSumCount(n1, n2, n3, n4); r != 1 {
		t.Fatalf("expect %d get %d", 1, r)
	}
}

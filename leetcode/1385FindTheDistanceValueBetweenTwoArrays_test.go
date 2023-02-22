package leetcode

import "testing"

func TestFindTheDistanceValue(t *testing.T) {
	a1, a2 := []int{4, 5, 8}, []int{10, 9, 1, 8}
	if r := findTheDistanceValue(a1, a2, 2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

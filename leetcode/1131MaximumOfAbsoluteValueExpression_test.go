package leetcode

import "testing"

func TestMaxAbsVAlExpr(t *testing.T) {
	arr1, arr2 := []int{1, 2, 3, 4}, []int{-1, 4, 5, 6}
	if r := maxAbsValExpr(arr1, arr2); r != 13 {
		t.Fatalf("expect 13 get  %d", r)
	}

	arr1, arr2 = []int{1, -2, -5, 0, 10}, []int{0, -2, -1, -7, -4}
	if r := maxAbsValExpr(arr1, arr2); r != 20 {
		t.Fatalf("expect 20 get %d", r)
	}
}

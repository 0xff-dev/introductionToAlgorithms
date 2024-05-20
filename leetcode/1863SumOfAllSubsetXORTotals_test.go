package leetcode

import "testing"

func TestSubsetXORSum(t *testing.T) {
	nums := []int{1, 3}
	exp := 6
	if r := subsetXORSum(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	nums = []int{5, 1, 6}
	exp = 28
	if r := subsetXORSum(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{3, 4, 5, 6, 7, 8}
	exp = 480
	if r := subsetXORSum(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

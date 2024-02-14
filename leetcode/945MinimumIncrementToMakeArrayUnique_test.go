package leetcode

import "testing"

func TestMinIncrementForUnique(t *testing.T) {
	nums := []int{1, 2, 2}
	exp := 1
	if r := minIncrementForUnique(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{3, 2, 1, 2, 1, 7}
	exp = 6
	if r := minIncrementForUnique(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{1, 2, 3, 3, 3, 3, 2, 4, 4, 4, 45, 5}
	exp = 32
	if r := minIncrementForUnique(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

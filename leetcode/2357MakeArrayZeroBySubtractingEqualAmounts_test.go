package leetcode

import "testing"

func TestMinimumOperations(t *testing.T) {
	nums := []int{1, 3, 0, 5, 0}
	if r := minimumOperations(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	nums = []int{0}
	if r := minimumOperations(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	nums = []int{1}
	if r := minimumOperations(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

}

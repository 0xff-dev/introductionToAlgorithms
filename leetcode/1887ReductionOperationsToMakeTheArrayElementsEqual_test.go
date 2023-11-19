package leetcode

import "testing"

func TestReductionOperations(t *testing.T) {
	nums := []int{5, 1, 3}
	if r := reductionOperations(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	nums = []int{1, 1, 1}
	if r := reductionOperations(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	nums = []int{1, 1, 2, 2, 3}
	if r := reductionOperations(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

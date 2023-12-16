package leetcode

import "testing"

func TestFindMaximumXOR(t *testing.T) {
	nums := []int{3, 10, 5, 25, 2, 8}
	if r := findMaximumXOR(nums); r != 28 {
		t.Fatalf("expect 28 get %d", r)
	}
	nums = []int{0}
	if r := findMaximumXOR(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	nums = []int{4, 6, 7}
	if r := findMaximumXOR(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

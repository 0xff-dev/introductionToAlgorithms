package leetcode

import "testing"

func TestWiggleMaxLength(t *testing.T) {
	nums := []int{1, 7, 4, 9, 2, 5}
	if r := wiggleMaxLength(nums); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	nums = []int{1}
	if r := wiggleMaxLength(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	nums = []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	if r := wiggleMaxLength(nums); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}

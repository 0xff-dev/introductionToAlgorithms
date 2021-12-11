package leetcode

import "testing"

func TestSingleNumber1(t *testing.T) {
	nums := []int{2, 2, 1}
	if r := singleNumber1(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	nums = []int{4, 1, 2, 1, 2}
	if r := singleNumber1(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	nums = []int{1}
	if r := singleNumber1(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

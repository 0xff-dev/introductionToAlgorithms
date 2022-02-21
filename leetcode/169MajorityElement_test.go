package leetcode

import "testing"

func TestMajorityElement(t *testing.T) {
	nums := []int{3, 2, 3}
	if r := majorityElement(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	if r := majorityElement(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{1, 2, 1, 2, 2, 1, 1}
	if r := majorityElement(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	nums = []int{1, 2, 1, 2, 2, 1, 2}
	if r := majorityElement(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{1, 1, 2, 2, 3, 2, 2, 1, 1, 1, 3}
	if r := majorityElement(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

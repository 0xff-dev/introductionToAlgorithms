package leetcode

import "testing"

func TestSpecialArray(t *testing.T) {
	nums := []int{3, 5}
	if r := specialArray(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{0, 0}
	if r := specialArray(nums); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	nums = []int{0, 4, 3, 0, 4}
	if r := specialArray(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

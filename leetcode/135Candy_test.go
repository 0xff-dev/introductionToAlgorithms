package leetcode

import "testing"

func TestCandy(t *testing.T) {
	nums := []int{1, 0, 2}
	if r := candy(nums); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	nums = []int{12, 4, 3, 11, 34, 34, 1, 67}
	if r := candy(nums); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}

	nums = []int{1}
	if r := candy(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	nums = []int{1, 2}
	if r := candy(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

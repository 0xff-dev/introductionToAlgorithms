package leetcode

import "testing"

func TestBinarySea(t *testing.T) {
	nums, target := []int{-1, 0, 3, 5, 9, 12}, 9
	if idx := search(nums, target); idx != 4 {
		t.Fatalf("expect 4 get %d", idx)
	}

	target = 2
	if idx := search(nums, target); idx != -1 {
		t.Fatalf("expect -1 get %d", idx)
	}

	nums, target = []int{1}, 1
	if idx := search(nums, target); idx != 0 {
		t.Fatalf("expect 0 get %d", idx)
	}

	nums, target = []int{}, 1
	if idx := search(nums, target); idx != -1 {
		t.Fatalf("expect -1 get %d", idx)
	}
}

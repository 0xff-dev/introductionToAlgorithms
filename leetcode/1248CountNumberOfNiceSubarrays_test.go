package leetcode

import "testing"

func TestNumberOfSubarrays(t *testing.T) {
	nums, k := []int{1, 1, 2, 1, 1}, 3
	if r := numberOfSubarrays(nums, k); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums, k = []int{2, 4, 6}, 1
	if r := numberOfSubarrays(nums, k); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	nums, k = []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2
	if r := numberOfSubarrays(nums, k); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}

	nums, k = []int{91473, 45388, 24720, 35841, 29648, 77363, 86290, 58032, 53752, 87188, 34428, 85343, 19801, 73201}, 4
	if r := numberOfSubarrays(nums, k); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	nums, k = []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 1
	if r := numberOfSubarrays(nums, k); r != 24 {
		t.Fatalf("expect 20 get %d", r)
	}

	nums, k = []int{2,1,1}, 1
	if r := numberOfSubarrays(nums, k); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

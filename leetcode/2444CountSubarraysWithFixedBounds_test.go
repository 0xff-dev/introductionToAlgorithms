package leetcode

import "testing"

func TestCountSubarrays(t *testing.T) {
	nums := []int{35054, 398719, 945315, 945315, 820417, 945315, 35054, 945315, 171832, 945315, 35054, 109750, 790964, 441974, 552913}
	minK, maxK := 35054, 945315
	if r := countSubarrays(nums, minK, maxK); r != 81 {
		t.Fatalf("expect 81 get %d", r)
	}
	nums = []int{1, 3, 5, 2, 7, 5}
	minK, maxK = 1, 5
	if r := countSubarrays(nums, minK, maxK); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{1, 1, 1, 1}
	minK, maxK = 1, 1
	if r := countSubarrays(nums, minK, maxK); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}

	nums = []int{978650, 978650, 978650, 68071, 52201, 68071, 186141, 978650, 978650, 267135, 68071, 717241, 242895, 68071, 582505, 978650, 68071, 68071}
	minK, maxK = 68071, 978650
	if r := countSubarrays(nums, minK, maxK); r != 57 {
		t.Fatalf("expect 57 get %d", r)
	}

	nums = []int{1, 4, 10, 2, 4, 1, 2, 10, 2, 4, 1, 2, 4, 1, 21, 43, 12}
	minK, maxK = 3, 13
	if r := countSubarrays(nums, minK, maxK); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

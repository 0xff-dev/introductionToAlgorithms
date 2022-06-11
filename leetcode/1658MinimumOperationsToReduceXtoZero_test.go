package leetcode

import "testing"

func TestMinOperations(t *testing.T) {
	nums, x := []int{1, 1, 4, 2, 3}, 5
	if r := minOperations(nums, x); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums, x = []int{5, 6, 7, 8, 9}, 4
	if r := minOperations(nums, x); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	nums, x = []int{3, 2, 20, 1, 1, 3}, 10
	if r := minOperations(nums, x); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	nums, x = []int{1, 1}, 3
	if r := minOperations(nums, x); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	nums, x = []int{10, 1, 1, 1, 1, 1}, 5
	if r := minOperations(nums, x); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	nums, x = []int{5, 2, 3, 1, 1}, 5
	if r := minOperations(nums, x); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

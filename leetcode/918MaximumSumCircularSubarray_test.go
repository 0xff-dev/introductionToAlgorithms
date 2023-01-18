package leetcode

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	/*
		nums := []int{1, -2, 3, -2}
		if r := maxSubarraySumCircular(nums); r != 3 {
			t.Fatalf("expect 3 get %d", r)
		}

		nums = []int{-3, -2, -3}
		if r := maxSubarraySumCircular(nums); r != -2 {
			t.Fatalf("expect -2 get %d", r)
		}

		nums = []int{5, -3, 5}
		if r := maxSubarraySumCircular(nums); r != 10 {
			t.Fatalf("expect 10 get %d", r)
		}

		nums = []int{1}
		if r := maxSubarraySumCircular(nums); r != 1 {
			t.Fatalf("expect 1 get %d", r)
		}

		nums = []int{1, 2, 3, 4}
		if r := maxSubarraySumCircular(nums); r != 10 {
			t.Fatalf("expect 10 get %d", r)
		}

		nums = []int{1, 2}
		if r := maxSubarraySumCircular(nums); r != 3 {
			t.Fatalf("expect 3 get %d", r)
		}

		nums = []int{1, -2, 3}
		if r := maxSubarraySumCircular(nums); r != 4 {
			t.Fatalf("expect 4 get %d", r)
		}

		nums = []int{-2,2,-2,9}
		if r := maxSubarraySumCircular(nums); r != 9 {
			t.Fatalf("expect 9 get %d", r)
		}
	*/

	nums := []int{8, -15, -29, -19}
	if r := maxSubarraySumCircular(nums); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}

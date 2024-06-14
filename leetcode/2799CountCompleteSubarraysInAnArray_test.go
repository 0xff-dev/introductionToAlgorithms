package leetcode

import "testing"

func TestCountCompleteSubarrays(t *testing.T) {
	nums := []int{1, 3, 1, 2, 2}
	exp := 4
	if r := countCompleteSubarrays(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	nums = []int{5, 5, 5, 5}
	exp = 10
	if r := countCompleteSubarrays(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

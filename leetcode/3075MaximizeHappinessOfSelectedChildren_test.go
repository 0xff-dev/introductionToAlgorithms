package leetcode

import "testing"

func TestMaximumHappinessSum(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 2
	exp := int64(4)
	if r := maximumHappinessSum(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	nums = []int{1, 1, 1, 1}
	k = 2
	exp = 1
	if r := maximumHappinessSum(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums = []int{2, 3, 4, 5}
	k = 1
	exp = 5
	if r := maximumHappinessSum(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

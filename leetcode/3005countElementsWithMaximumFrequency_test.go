package leetcode

import "testing"

func TestMaxFrequencyElements(t *testing.T) {
	nums := []int{1, 2, 2, 3, 1, 4}
	exp := 4
	if r := maxFrequencyElements(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	nums = []int{1, 2, 3, 4, 5}
	exp = 5
	if r := maxFrequencyElements(nums); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

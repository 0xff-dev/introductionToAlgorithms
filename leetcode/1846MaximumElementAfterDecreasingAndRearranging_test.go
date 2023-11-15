package leetcode

import "testing"

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	arr := []int{2, 2, 1, 2, 1}
	if r := maximumElementAfterDecrementingAndRearranging(arr); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	arr = []int{100, 1, 1000}
	if r := maximumElementAfterDecrementingAndRearranging(arr); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	arr = []int{1, 2, 3, 4, 5}
	if r := maximumElementAfterDecrementingAndRearranging(arr); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}

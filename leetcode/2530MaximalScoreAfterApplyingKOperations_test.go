package leetcode

import "testing"

func TestMaxKelements(t *testing.T) {
	nums := []int{10, 10, 10, 10, 10}
	n := 5
	if r := maxKelements(nums, n); r != 50 {
		t.Fatalf("expect 50 get %d", r)
	}
	nums = []int{1, 10, 3, 3, 3}
	n = 3
	if r := maxKelements(nums, n); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}
}

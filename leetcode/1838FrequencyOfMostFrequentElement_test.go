package leetcode

import "testing"

func TestMaxFrequency(t *testing.T) {
	nums := []int{9940, 9995, 9944, 9937, 9941, 9952, 9907, 9952, 9987, 9964, 9940, 9914, 9941, 9933, 9912, 9934, 9980, 9907, 9980, 9944, 9910, 9997}
	k := 7925
	if r := maxFrequency(nums, k); r != 22 {
		t.Fatalf("expect 22 get %d", r)
	}
	nums = []int{1000}
	k = 100
	if r := maxFrequency(nums, k); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

package leetcode

import "testing"

func TestTotalHammingDistance(t *testing.T) {
	nums := []int{4, 14, 2}
	if r := totalHammingDistance(nums); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	nums = []int{4, 14, 4}
	if r := totalHammingDistance(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

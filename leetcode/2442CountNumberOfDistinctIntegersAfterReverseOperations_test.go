package leetcode

import "testing"

func TestCountDistinctIntegers(t *testing.T) {
	nums := []int{1, 13, 10, 12, 31}
	if r := countDistinctIntegers(nums); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}

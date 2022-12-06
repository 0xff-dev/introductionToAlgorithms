package leetcode

import "testing"

func TestCountHillValley(t *testing.T) {
	nums := []int{2, 4, 1, 1, 6, 5}
	if r := countHillValley(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	nums = []int{6, 6, 5, 5, 4, 1}
	if r := countHillValley(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

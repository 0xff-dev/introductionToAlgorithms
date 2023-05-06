package leetcode

import "testing"

func TestNumSubseq(t *testing.T) {
	nums := []int{3, 5, 6, 7}
	target := 9
	if r := numSubseq(nums, target); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

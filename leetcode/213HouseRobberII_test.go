package leetcode

import "testing"

func TestRob1(t *testing.T) {
	nums := []int{2, 3, 2}
	if r := rob1(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	nums = []int{1, 2, 3, 1}
	if r := rob1(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	nums = []int{0}
	if r := rob1(nums); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

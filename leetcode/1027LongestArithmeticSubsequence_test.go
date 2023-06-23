package leetcode

import "testing"

func TestLongestArithSeqLength(t *testing.T) {
	nums := []int{3, 6, 9, 12}
	if r := longestArithSeqLength(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	nums = []int{9, 4, 7, 2, 10}
	if r := longestArithSeqLength(nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	nums = []int{20, 1, 15, 3, 10, 5, 8}
	if r := longestArithSeqLength(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	nums = []int{24, 13, 1, 100, 0, 94, 3, 0, 3}
	if r := longestArithSeqLength(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

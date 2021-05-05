package leetcode

import "testing"

func TestLengthOfLIS(t *testing.T) {
	input := []int{10, 9, 2, 5, 3, 7, 101, 18}
	if r := lengthOfLIS(input); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	input = []int{0, 1, 0, 3, 2, 3}
	if r := lengthOfLIS(input); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	input = []int{7, 7, 7, 7, 7}
	if r := lengthOfLIS(input); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	input = []int{0}
	if r := lengthOfLIS(input); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

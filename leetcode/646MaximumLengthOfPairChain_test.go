package leetcode

import "testing"

func TestFindLongestChain(t *testing.T) {
	pairs := [][]int{
		{1, 2}, {2, 3}, {3, 4},
	}
	if r := findLongestChain(pairs); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	pairs = [][]int{
		{1, 2}, {7, 8}, {4, 5},
	}
	if r := findLongestChain(pairs); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

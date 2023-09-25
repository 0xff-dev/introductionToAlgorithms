package leetcode

import "testing"

func TestMaxChunksToSorted(t *testing.T) {
	arr := []int{4, 3, 2, 1, 0}
	if r := maxChunksToSorted(arr); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	arr = []int{1, 0, 2, 3, 4}
	if r := maxChunksToSorted(arr); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

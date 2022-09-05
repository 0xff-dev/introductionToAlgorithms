package leetcode

import "testing"

func TestLastStoneWeight(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	if r := lastStoneWeight(stones); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	stones = []int{1, 1}
	if r := lastStoneWeight(stones); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	stones = []int{1}
	if r := lastStoneWeight(stones); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
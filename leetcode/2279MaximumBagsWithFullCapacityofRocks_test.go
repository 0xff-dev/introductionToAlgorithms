package leetcode

import "testing"

func TestMaximumBags(t *testing.T) {
	capacity, rocks, additionalRocks := []int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2
	if r := maximumBags(capacity, rocks, additionalRocks); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	capacity, rocks, additionalRocks = []int{10, 2, 2}, []int{2, 2, 0}, 100
	if r := maximumBags(capacity, rocks, additionalRocks); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

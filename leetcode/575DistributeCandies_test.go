package leetcode

import "testing"

func TestDistribute(t *testing.T) {
	candyType := []int{1, 1, 2, 2, 3, 3}
	if r := distributeCandies(candyType); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	candyType = []int{1, 1, 2, 3}
	if r := distributeCandies(candyType); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	candyType = []int{6, 6, 6, 6}
	if r := distributeCandies(candyType); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

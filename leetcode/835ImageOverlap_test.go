package leetcode

import "testing"

func TestLargestOverlap(t *testing.T) {
	i1, i2 := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}, [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}
	if r := largestOverlap(i1, i2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	i1, i2 = [][]int{{1}}, [][]int{{1}}
	if r := largestOverlap(i1, i2); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	i1, i2 = [][]int{{0}}, [][]int{{0}}
	if r := largestOverlap(i1, i2); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

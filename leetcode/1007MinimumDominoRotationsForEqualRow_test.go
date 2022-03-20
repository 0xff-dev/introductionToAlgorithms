package leetcode

import "testing"

func TestMinDominoRotations(t *testing.T) {
	tops, bottoms := []int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}
	if r := minDominoRotations(tops, bottoms); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	tops, bottoms = []int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}
	if r := minDominoRotations(tops, bottoms); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}

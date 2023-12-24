package leetcode

import "testing"

func TestNumberOfBoomerangs(t *testing.T) {
	points := [][]int{
		{0, 0}, {1, 0}, {2, 0},
	}
	if r := numberOfBoomerangs(points); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

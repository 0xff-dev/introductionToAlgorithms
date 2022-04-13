package leetcode

import "testing"

func TestFindRotation(t *testing.T) {
	mat, target := [][]int{
		{0, 1},
		{1, 0},
	}, [][]int{
		{1, 0},
		{0, 1},
	}
	if !findRotation(mat, target) {
		t.Fatalf("expect true get false")
	}

	mat, target = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}, [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 0, 0},
	}
	if !findRotation(mat, target) {
		t.Fatalf("expect true get false")
	}
}

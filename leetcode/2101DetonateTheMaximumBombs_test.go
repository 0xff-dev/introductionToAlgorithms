package leetcode

import "testing"

func TestMaximumDetonation(t *testing.T) {
	bombs := [][]int{
		{2, 1, 3}, {6, 1, 4},
	}
	if r := maximumDetonation(bombs); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	bombs = [][]int{
		{1, 1, 5}, {10, 10, 5},
	}
	if r := maximumDetonation(bombs); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	bombs = [][]int{
		{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4},
	}
	if r := maximumDetonation(bombs); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}

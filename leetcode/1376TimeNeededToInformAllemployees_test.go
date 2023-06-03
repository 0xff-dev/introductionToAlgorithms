package leetcode

import "testing"

func TestNumOfMinutes(t *testing.T) {
	n := 1
	headID := 0
	manager := []int{-1}
	infoTime := []int{0}
	if r := numOfMinutes(n, headID, manager, infoTime); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	n = 6
	headID = 2
	manager = []int{2, 2, -1, 2, 2, 2}
	infoTime = []int{0, 0, 1, 0, 0, 0}
	if r := numOfMinutes(n, headID, manager, infoTime); r != 1 {
		t.Fatalf("expect %d get %d", 1, r)
	}

	n = 11
	headID = 4
	manager = []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}
	infoTime = []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}
	if r := numOfMinutes(n, headID, manager, infoTime); r != 2560 {
		t.Fatalf("expect 2560 get %d", r)
	}
}

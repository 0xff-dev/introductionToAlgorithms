package leetcode

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	rooms := [][]int{
		{1},
		{2},
		{3},
		{},
	}

	if !canVisitAllRooms(rooms) {
		t.Fatalf("expect true get false")
	}

	rooms = [][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}

	if canVisitAllRooms(rooms) {
		t.Fatalf("expect false get true")
	}
}

package leetcode

import "testing"

func TestCarPooling(t *testing.T) {
	trips := [][]int{
		{2, 1, 5},
		{3, 3, 7},
	}
	if carPooling(trips, 4) {
		t.Fatalf("expect false get true")
	}

	if !carPooling(trips, 5) {
		t.Fatalf("expect true get false")
	}

	trips = [][]int{
		{1, 1, 8},
		{2, 2, 7},
		{3, 3, 6},
	}

	if carPooling(trips, 5) {
		t.Fatalf("expect false get true")
	}

	if !carPooling(trips, 6) {
		t.Fatalf("expect true get false")
	}

	trips = [][]int{
		{2, 1, 5},
		{3, 5, 7},
	}
	if !carPooling(trips, 3) {
		t.Fatalf("expect true get false")
	}
}

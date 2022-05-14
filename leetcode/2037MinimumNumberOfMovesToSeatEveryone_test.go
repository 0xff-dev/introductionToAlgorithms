package leetcode

import "testing"

func TestMinMovesToSeat(t *testing.T) {
	seats, students := []int{3, 1, 5}, []int{2, 7, 4}
	if r := minMovesToSeat(seats, students); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	seats, students = []int{4, 1, 5, 9}, []int{1, 3, 2, 6}
	if r := minMovesToSeat(seats, students); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	seats, students = []int{2, 2, 6, 6}, []int{1, 3, 2, 6}
	if r := minMovesToSeat(seats, students); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

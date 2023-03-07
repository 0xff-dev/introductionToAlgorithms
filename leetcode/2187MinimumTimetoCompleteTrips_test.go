package leetcode

import "testing"

func TestMinimumTime(t *testing.T) {
	time := []int{1, 2, 3}
	if r := minimumTime(time, 5); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	time = []int{2}
	if r := minimumTime(time, 1); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

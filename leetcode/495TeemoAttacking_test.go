package leetcode

import "testing"

func TestFindPoisonedDuration(t *testing.T) {
	timeSeries := []int{1, 4}
	if r := findPoisonedDuration(timeSeries, 2); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	timeSeries = []int{1, 2}
	if r := findPoisonedDuration(timeSeries, 2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	timeSeries = []int{1}
	if r := findPoisonedDuration(timeSeries, 2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	timeSeries = []int{1, 2, 3, 4, 5}
	if r := findPoisonedDuration(timeSeries, 5); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
}

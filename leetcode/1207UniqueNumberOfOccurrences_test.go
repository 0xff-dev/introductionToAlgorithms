package leetcode

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	arr := []int{1, 2, 2, 1, 1, 3}
	if !uniqueOccurrences(arr) {
		t.Fatalf("expect true get false")
	}

	arr = []int{1, 2}
	if uniqueOccurrences(arr) {
		t.Fatalf("expect false get true")
	}

	arr = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	if !uniqueOccurrences(arr) {
		t.Fatalf("expect true get false")
	}
}

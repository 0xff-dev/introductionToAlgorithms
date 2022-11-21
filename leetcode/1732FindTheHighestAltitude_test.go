package leetcode

import "testing"

func TestLargestAltitude(t *testing.T) {
	gain := []int{-5, 1, 5, 0, -7}
	if r := largestAltitude(gain); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	gain = []int{-4, -3, -2, -1, 4, 3, 2}
	if r := largestAltitude(gain); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

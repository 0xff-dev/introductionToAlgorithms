package leetcode

import "testing"

func TestFindMinDifference(t *testing.T) {
	tp := []string{"23:59", "00:00"}
	if r := findMinDifference(tp); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	tp = []string{"01:39", "10:26", "21:43"}
	if r := findMinDifference(tp); r != 236 {
		t.Fatalf("expect 236 get %d", r)
	}
}

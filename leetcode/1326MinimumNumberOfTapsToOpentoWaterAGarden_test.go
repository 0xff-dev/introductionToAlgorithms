package leetcode

import "testing"

func TestMinTaps(t *testing.T) {
	n := 5
	ranges := []int{3, 4, 1, 1, 0, 0}
	if r := minTaps(n, ranges); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	n = 3
	ranges = []int{0, 0, 0, 0}
	if r := minTaps(n, ranges); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}

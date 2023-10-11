package leetcode

import "testing"

func TestSumDistance(t *testing.T) {
	n := []int{-2, 0, 2}
	s := "RLL"
	d := 3
	if r := sumDistance(n, s, d); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	n = []int{1, -67, 68, -26, -13, -40, -56, 62, 21}
	s = "LLLRLLRRR"
	d = 4
	if r := sumDistance(n, s, d); r != 2106 {
		t.Fatalf("expect 2106 get %d", r)
	}
}

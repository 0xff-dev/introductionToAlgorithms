package leetcode

import "testing"

func TestMaxSatisfaction(t *testing.T) {
	a := []int{-1, -8, 0, 5, -9}
	if r := maxSatisfaction(a); r != 14 {
		t.Fatalf("expect 14 get %d", r)
	}
	a = []int{4, 3, 2}
	if r := maxSatisfaction(a); r != 20 {
		t.Fatalf("expect 20 get %d", r)
	}

	a = []int{-1, -4, -5}
	if r := maxSatisfaction(a); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	a = []int{-2, 5, -1, 0, 3, -3}
	if r := maxSatisfaction(a); r != 35 {
		t.Fatalf("expect 35 get %d", r)
	}
}

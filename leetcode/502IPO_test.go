package leetcode

import "testing"

func TestFindmaximizeCapital(t *testing.T) {
	k, w := 2, 0
	p, c := []int{1, 2, 3}, []int{0, 1, 1}
	if r := findMaximizedCapital(k, w, p, c); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	k, w = 3, 0
	p, c = []int{1, 2, 3}, []int{0, 1, 2}
	if r := findMaximizedCapital(k, w, p, c); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}

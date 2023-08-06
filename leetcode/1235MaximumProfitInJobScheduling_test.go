package leetcode

import "testing"

func TestJobScheduling(t *testing.T) {
	s, e, p := []int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}
	if r := jobScheduling(s, e, p); r != 120 {
		t.Fatalf("expect 120 get %d", r)
	}

	s, e, p = []int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}
	if r := jobScheduling(s, e, p); r != 150 {
		t.Fatalf("expect 150 get %d", r)
	}
	s, e, p = []int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}
	if r := jobScheduling(s, e, p); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}

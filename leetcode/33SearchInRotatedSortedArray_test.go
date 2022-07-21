package leetcode

import "testing"

func TestSearch(t *testing.T) {
	n := []int{4, 5, 6, 7, 0, 1, 2}
	if r := search33(n, 0); r != 4 {
		t.Fatalf("expect %d get %d", 4, r)
	}
	n = []int{2, 3, 4, 5, 1}
	if r := search33(n, 1); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	n = []int{4, 5, 6, 7, 0, 1, 2}
	if r := search33(n, 3); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	n = []int{3, 1}
	if r := search33(n, 3); r != 0 {
		t.Fatalf("expect 1 get %d", r)
	}
}

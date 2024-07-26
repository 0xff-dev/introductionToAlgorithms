package leetcode

import "testing"

func TestFindTheCity(t *testing.T) {
	n, e, d, exp := 4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4, 3
	if r := findTheCity(n, e, d); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, e, d, exp = 5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2, 0
	if r := findTheCity(n, e, d); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, e, d, exp = 6, [][]int{{0, 1, 10}, {0, 2, 1}, {2, 3, 1}, {1, 3, 1}, {1, 4, 1}, {4, 5, 10}}, 20, 5
	if r := findTheCity(n, e, d); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

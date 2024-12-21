package leetcode

import "testing"

func TestMaxKDivisibleCompnents(t *testing.T) {
	n, edges, values, k, exp := 5, [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}, []int{1, 8, 1, 4, 4}, 6, 2
	if r := maxKDivisibleComponents(n, edges, values, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, edges, values, k, exp = 7, [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}, []int{3, 0, 6, 1, 5, 2, 1}, 3, 3
	if r := maxKDivisibleComponents(n, edges, values, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

package leetcode

import "testing"

func TestReachableNodes(t *testing.T) {
	n, edges, restricted, exp := 7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5}, 4
	if r := reachableNodes(n, edges, restricted); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, edges, restricted, exp = 7, [][]int{{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}}, []int{4, 2, 1}, 3
	if r := reachableNodes(n, edges, restricted); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

package leetcode

import "testing"

func TestMaxTotalFruits(t *testing.T) {
	fruits, startPos, k, exp := [][]int{{2, 8}, {6, 3}, {8, 6}}, 5, 4, 9
	if r := maxTotalFruits(fruits, startPos, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	fruits, startPos, k, exp = [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, 5, 4, 14
	if r := maxTotalFruits(fruits, startPos, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	fruits, startPos, k, exp = [][]int{{0, 3}, {6, 4}, {8, 5}}, 3, 2, 0
	if r := maxTotalFruits(fruits, startPos, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

package leetcode

import "testing"

func TestMaxSumSubmatrix(t *testing.T) {
	m := [][]int{
		{1, 0, 1}, {0, -2, 3},
	}
	if r := maxSumSubmatrix(m, 2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	m = [][]int{{2, 2, -1}}
	if r := maxSumSubmatrix(m, 3); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

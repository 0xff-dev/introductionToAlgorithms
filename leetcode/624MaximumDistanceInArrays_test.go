package leetcode

import "testing"

func TestMaxDistance624(t *testing.T) {
	arrays, exp := [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}, 4
	if r := maxDistance624(arrays); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	arrays, exp = [][]int{{1}, {1}}, 0
	if r := maxDistance624(arrays); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

package leetcode

import "testing"

func TestKthLargestValue(t *testing.T) {
	matrix := [][]int{
		{5, 2}, {1, 6},
	}
	k := 1
	exp := 7
	if r := kthLargestValue(matrix, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	k = 2
	exp = 5
	if r := kthLargestValue(matrix, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	k = 3
	exp = 4
	if r := kthLargestValue(matrix, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

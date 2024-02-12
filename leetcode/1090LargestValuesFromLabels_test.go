package leetcode

import "testing"

func TestLargestValsFromLabels(t *testing.T) {

	values := []int{5, 4, 3, 2, 1}
	labels := []int{1, 1, 2, 2, 3}
	want := 3
	limit := 1
	exp := 9
	if r := largestValsFromLabels(values, labels, want, limit); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	values = []int{2, 1, 5, 2, 8}
	labels = []int{2, 0, 2, 2, 2}
	want = 3
	limit = 2
	exp = 14
	if r := largestValsFromLabels(values, labels, want, limit); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

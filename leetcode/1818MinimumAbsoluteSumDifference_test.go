package leetcode

import "testing"

func TestMinAbsoluteSumDiff(t *testing.T) {
	n1, n2 := []int{1, 7, 5}, []int{2, 3, 5}
	if r := minAbsoluteSumDiff(n1, n2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	n1, n2 = []int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}
	if r := minAbsoluteSumDiff(n1, n2); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	n1, n2 = []int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}
	if r := minAbsoluteSumDiff(n1, n2); r != 20 {
		t.Fatalf("expect 20 get %d", r)
	}

}

package leetcode

import "testing"

func TestMinDays1482(t *testing.T) {
	b, m, k := []int{1, 10, 3, 10, 2}, 3, 1
	exp := 3
	if r := minDays1482(b, m, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	b, m, k = []int{1, 10, 3, 10, 2}, 3, 2
	exp = -1
	if r := minDays1482(b, m, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	b, m, k = []int{7, 7, 7, 7, 12, 7, 7}, 2, 3
	exp = 12
	if r := minDays1482(b, m, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

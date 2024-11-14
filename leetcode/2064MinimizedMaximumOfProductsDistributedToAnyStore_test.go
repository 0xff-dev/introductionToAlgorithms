package leetcode

import "testing"

func TestMinimizedMaximum(t *testing.T) {
	n, quantities, exp := 6, []int{11, 6}, 3
	if r := minimizedMaximum(n, quantities); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, quantities, exp = 7, []int{15, 10, 10}, 5
	if r := minimizedMaximum(n, quantities); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, quantities, exp = 1, []int{100000}, 100000
	if r := minimizedMaximum(n, quantities); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

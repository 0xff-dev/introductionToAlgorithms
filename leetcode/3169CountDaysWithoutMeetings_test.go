package leetcode

import "testing"

func TestCountDays(t *testing.T) {
	days, meetings, exp := 10, [][]int{{5, 7}, {1, 3}, {9, 10}}, 2
	if r := countDays(days, meetings); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	days, meetings, exp = 5, [][]int{{2, 4}, {1, 3}}, 1
	if r := countDays(days, meetings); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	days, meetings, exp = 6, [][]int{{1, 6}}, 0
	if r := countDays(days, meetings); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

}

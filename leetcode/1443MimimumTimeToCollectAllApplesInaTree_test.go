package leetcode

import "testing"

func TestMinTime(t *testing.T) {
	n, edges, hasApple := 7, [][]int{
		{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6},
	}, []bool{false, false, true, false, false, true, false}
	if r := minTime(n, edges, hasApple); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	hasApple = []bool{false, false, true, false, true, true, false}
	if r := minTime(n, edges, hasApple); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}

	n, edges, hasApple = 1, [][]int{}, []bool{true}
	if r := minTime(n, edges, hasApple); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	n, edges, hasApple = 4, [][]int{{0, 1}, {1, 2}, {0, 3}}, []bool{true, true, true, true}
	if r := minTime(n, edges, hasApple); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	n, edges, hasApple = 4, [][]int{{0, 2}, {0, 3}, {1, 2}}, []bool{false, true, false, false}
	if r := minTime(n, edges, hasApple); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

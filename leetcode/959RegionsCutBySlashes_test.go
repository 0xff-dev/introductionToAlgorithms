package leetcode

import "testing"

func TestRegionsBySlashes(t *testing.T) {
	grid, exp := []string{" /", "/ "}, 2
	if r := regionsBySlashes(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	grid, exp = []string{"/\\", "\\/"}, 5
	if r := regionsBySlashes(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	grid, exp = []string{"\\\\\\", "\\\\/", " //"}, 5
	if r := regionsBySlashes(grid); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

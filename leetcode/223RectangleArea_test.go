package leetcode

import "testing"

func TestComputeArea(t *testing.T) {
	ax1, ay1, ax2, ay2 := -3, 0, 3, 4
	bx1, by1, bx2, by2 := 0, -1, 9, 2
	if r := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2); r != 45 {
		t.Fatalf("expect 45 get %d", r)
	}

	ax1, ay1, ax2, ay2 = -2, -2, 2, 2
	bx1, by1, bx2, by2 = 3, 3, 4, 4
	if r := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}
	ax1, ay1, ax2, ay2 = -5, -5, 0, -4
	bx1, by1, bx2, by2 = -3, -3, 3, 3
	if r := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2); r != 41 {
		t.Fatalf("expect 41 get %d", r)
	}
}

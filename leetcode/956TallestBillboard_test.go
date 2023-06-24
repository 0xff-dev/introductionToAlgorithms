package leetcode

import "testing"

func TestTallestBillboard(t *testing.T) {
	rods := []int{1, 2, 3, 6}
	if r := tallestBillboard(rods); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	rods = []int{1, 2, 3, 4, 5, 6}
	if r := tallestBillboard(rods); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	rods = []int{1, 2}
	if r := tallestBillboard(rods); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	rods = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 22}
	//rods = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	if r := tallestBillboard(rods); r != 106 {
		t.Fatalf("expect 106 get %d", r)
	}
}

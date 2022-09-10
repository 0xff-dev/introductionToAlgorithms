package leetcode

import "testing"

func TestMaxProfit188(t *testing.T) {
	k := 2
	prices := []int{2, 4, 1}
	//if r := maxProfit188(k, prices); r != 2 {
	//	t.Fatalf("expect 2 get %d", r)
	//}
	//
	//prices = []int{3, 2, 6, 5, 0, 3}
	//if r := maxProfit188(k, prices); r != 7 {
	//	t.Fatalf("expect 7 get %d", r)
	//}

	prices = []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	if r := maxProfit188(k, prices); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}
}

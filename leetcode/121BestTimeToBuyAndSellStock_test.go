package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	if r := maxProfit(prices); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	if r := maxProfit2(prices); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	prices = []int{7, 6, 4, 3, 1}
	if r := maxProfit(prices); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	if r := maxProfit2(prices); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

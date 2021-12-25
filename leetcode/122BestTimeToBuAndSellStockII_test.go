package leetcode

import "testing"

func TestMaxProfitII(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	if r := maxProfitII(prices); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}

	prices = []int{1, 2, 3, 4, 5}
	if r := maxProfitII(prices); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	prices = []int{7, 6, 4, 3, 1}
	if r := maxProfitII(prices); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	prices = []int{3, 2, 6, 5, 0, 3}
	if r := maxProfitII(prices); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}

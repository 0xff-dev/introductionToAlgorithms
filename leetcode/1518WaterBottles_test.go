package leetcode

import "testing"

func TestNumWaterBottles(t *testing.T) {
	numBottles, numExchange := 9, 3
	if r := numWaterBottles(numBottles, numExchange); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}

	numBottles, numExchange = 15, 4
	if r := numWaterBottles(numBottles, numExchange); r != 19 {
		t.Fatalf("expect 19 get %d", r)
	}
}

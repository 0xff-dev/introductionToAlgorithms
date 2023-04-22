package leetcode

import (
	"testing"
)

func TestClosestCost(t *testing.T) {
	baseCosts := []int{1, 7}
	toppingCosts := []int{3, 4}
	target := 10
	if r := closestCost(baseCosts, toppingCosts, target); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	baseCosts = []int{2, 3}
	toppingCosts = []int{4, 5, 100}
	target = 18
	if r := closestCost(baseCosts, toppingCosts, target); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}

	baseCosts = []int{3, 10}
	toppingCosts = []int{2, 5}
	target = 9
	if r := closestCost(baseCosts, toppingCosts, target); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	baseCosts = []int{10}
	toppingCosts = []int{1}
	target = 1
	if r := closestCost(baseCosts, toppingCosts, target); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}

	baseCosts = []int{9}
	toppingCosts = []int{7}
	target = 10
	if r := closestCost(baseCosts, toppingCosts, target); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}

	baseCosts = []int{3}
	toppingCosts = []int{3}
	target = 9
	if r := closestCost(baseCosts, toppingCosts, target); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}

	baseCosts = []int{9, 10, 1}
	toppingCosts = []int{1, 8, 8, 1, 1, 8}
	target = 8
	if r := closestCost(baseCosts, toppingCosts, target); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}

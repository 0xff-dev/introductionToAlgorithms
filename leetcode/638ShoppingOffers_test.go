package leetcode

import "testing"

func TestShoppingOffers(t *testing.T) {
	price := []int{2, 5}
	special := [][]int{
		{3, 0, 5}, {1, 2, 10},
	}
	needs := []int{3, 2}
	exp := 14
	if r := shoppingOffers(price, special, needs); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	price = []int{2, 3, 4}
	special = [][]int{
		{1, 1, 0, 4}, {2, 2, 1, 9},
	}
	needs = []int{1, 2, 1}
	exp = 11
	if r := shoppingOffers(price, special, needs); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}

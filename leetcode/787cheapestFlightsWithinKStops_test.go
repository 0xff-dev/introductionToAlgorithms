package leetcode

import "testing"

func TestFindCheapestPrice(t *testing.T) {
	flights := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	n, src, dst, k := 4, 0, 3, 1
	if r := findCheapestPrice(n, flights, src, dst, k); r != 700 {
		t.Fatalf("expect 700 get %d", r)
	}
}

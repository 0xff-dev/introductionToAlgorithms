package leetcode

import "testing"

func TestMaximumWealth(t *testing.T) {
	accounts := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	if r := maximumWealth(accounts); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	accounts = [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}

	if r := maximumWealth(accounts); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}

	accounts = [][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
	}

	if r := maximumWealth(accounts); r != 17 {
		t.Fatalf("expect 17 get %d", r)
	}
}

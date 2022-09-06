package leetcode

import "testing"

func TestMaximumPopulation(t *testing.T) {
	logs := [][]int{
		{1993, 1999},
		{2000, 2010},
	}

	if r := maximumPopulation(logs); r != 1993 {
		t.Fatalf("expect 1993 get %d", r)
	}

	logs = [][]int{
		{1950, 1961},
		{1960, 1971},
		{1970, 1981},
	}
	if r := maximumPopulation(logs); r != 1960 {
		t.Fatalf("expect 1960 get %d", r)
	}
}

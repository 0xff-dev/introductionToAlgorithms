package leetcode

import "testing"

func TestSmallestStringWithSwaps(t *testing.T) {
	s := "dcab"
	pairs := [][]int{{0, 3}, {1, 2}}
	if r := smallestStringWithSwaps(s, pairs); r != "bacd" {
		t.Fatalf("expect bacd get %s", r)
	}

	s = "dcab"
	pairs = [][]int{
		{0, 3}, {1, 2}, {0, 2},
	}
	if r := smallestStringWithSwaps(s, pairs); r != "abcd" {
		t.Fatalf("expect abcd get %s", r)
	}

	s = "cba"
	pairs = [][]int{
		{0, 1}, {1, 2},
	}
	if r := smallestStringWithSwaps(s, pairs); r != "abc" {
		t.Fatalf("expect abc get %s", r)
	}
}

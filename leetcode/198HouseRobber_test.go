package leetcode

import "testing"

func TestRob(t *testing.T) {
	input := []int{1, 2, 3, 1}
	if r := rob(input); r != 4 {
		t.Fatalf("expect 4 get  %d", r)
	}
	input = []int{2, 7, 9, 3, 1}
	if r := rob(input); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}

	input = []int{2, 1, 1, 2}
	if r := rob(input); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	input = []int{2, 1}
	if r := rob(input); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	input = []int{1, 3, 1}
	if r := rob(input); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

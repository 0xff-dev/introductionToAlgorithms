package leetcode

import "testing"

func TestTrap(t *testing.T) {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	if r := trap(input); r != 6 {
		t.Fatal("should return 6 get ", r)
	}

	input = []int{4, 2, 0, 3, 2, 5}
	if r := trap(input); r != 9 {
		t.Fatal("should return 9 get ", r)
	}

	input = []int{1}
	if r := trap(input); r != 0 {
		t.Fatal("should return 0 get ", r)
	}

	input = []int{1, 0, 1}
	if r := trap(input); r != 1 {
		t.Fatal("should return 1 get ", r)
	}

	input = []int{1, 2, 3, 4}
	if r := trap(input); r != 0 {
		t.Fatal("should return 0 get ", r)
	}

	input = []int{1, 0}
	if r := trap(input); r != 0 {
		t.Fatal("should return 0 get ", r)
	}

	input = []int{0, 0, 1}
	if r := trap(input); r != 0 {
		t.Fatal("should return 0 get r")
	}

	input = []int{1, 0, 0}
	if r := trap(input); r != 0 {
		t.Fatal("should return 0 get ", r)
	}
}

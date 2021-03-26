package leetcode

import "testing"

func TestMaxProduct(t *testing.T) {
	input := []int{2, 3, -2, 4}
	if r := maxProduct(input); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	input = []int{2, 3, -2, -2}
	if r := maxProduct(input); r != 24 {
		t.Fatalf("expect 24 get %d", r)
	}

	input = []int{-1, -2, 2, -3, -1}
	if r := maxProduct(input); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}

	input = []int{-2, 0, -1}
	if r := maxProduct(input); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	input = []int{-2, 0, 1, 2}
	if r := maxProduct(input); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	input = []int{-1, 2, -3, -4, 0, 2}
	if r := maxProduct(input); r != 24 {
		t.Fatalf("expect 24 get %d", r)
	}

	input = []int{-2, 2, -3, -4, 0, 2}
	if r := maxProduct(input); r != 24 {
		t.Fatalf("expect 24 get %d", r)
	}


	input = []int{0, 0, 1, 2}
	if r := maxProduct(input); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	input = []int{0, 1, 0, 2}
	if r := maxProduct(input); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	input = []int{1, 2, 3, 4}
	if r := maxProduct(input); r != 24 {
		t.Fatalf("expect 24 get %d", r)
	}

	input = []int{-1, -2, 0, 2, -3, -4, 6, 0, 8, 9}
	if r := maxProduct(input); r != 144 {
		t.Fatalf("expect 144 get %d", r)
	}
}

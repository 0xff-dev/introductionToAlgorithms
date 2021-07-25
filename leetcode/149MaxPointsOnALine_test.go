package leetcode

import "testing"

func TestGCD(t *testing.T) {
	a, b := 2, 2
	if r := gcd(a, b); a != b && r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	a, b = 18, 4
	if r := gcd(a, b); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	a, b = 9, 5
	if r := gcd(a, b); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

func TestMaxPoints(t *testing.T) {
	points := [][]int{
		{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4},
	}
	if r := maxPoints(points); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	points = [][]int{
		{1, 1}, {2, 2}, {3, 3},
	}
	if r := maxPoints(points); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	points = [][]int{
		{4, 5}, {4, -1}, {4, 0},
	}

	if r := maxPoints(points); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	points = [][]int{
		{0, 0}, {4, 5}, {7, 8}, {8, 9}, {5, 6}, {3, 4}, {1, 1},
	}

	if r := maxPoints(points); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}

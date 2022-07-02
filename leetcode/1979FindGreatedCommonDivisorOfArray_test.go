package leetcode

import "testing"

func TestGCD1979(t *testing.T) {
	a, b := 16, 6
	if r := gcd1979(a, b); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	a, b = 8, 3
	if r := gcd1979(a, b); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

func TestFindGCD(t *testing.T) {
	nums := []int{2, 5, 6, 9, 10}
	if r := findGCD(nums); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{7, 5, 6, 8, 3}
	if r := findGCD(nums); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

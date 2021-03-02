package leetcode

import "testing"

func TestGetPermutation(t *testing.T) {
	n, k := 3, 3
	if r := getPermutation(n, k); r != "213" {
		t.Fatalf("expect 213 get %s", r)
	}

	n, k = 4, 9
	if r := getPermutation(n, k); r != "2314" {
		t.Fatalf("expect 2314 get %s", r)
	}

	n, k = 1, 1
	if r := getPermutation(n, k); r != "1" {
		t.Fatalf("expect 1 get %s", r)
	}
	n, k = 2, 1
	if r := getPermutation(n, k); r != "12" {
		t.Fatalf("expect 21 get %s", r)
	}
	n, k = 2, 2
	if r := getPermutation(n, k); r != "21" {
		t.Fatalf("expect 21 get %s", r)
	}

	n, k = 4, 24
	if r := getPermutation(n, k); r != "4321" {
		t.Fatalf("expect 4321 get %s", r)
	}
}

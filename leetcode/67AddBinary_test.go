package leetcode

import "testing"

func TestAddBinary(t *testing.T) {
	a, b := "11", "1"
	if r := addBinary(a, b); r != "100" {
		t.Fatalf("expect 100 get %s", r)
	}

	a, b = "1010", "1011"
	if r := addBinary(a, b); r != "10101" {
		t.Fatalf("expect 10101 get %s", r)
	}

	a, b = "11", "11"
	if r := addBinary(a, b); r != "110" {
		t.Fatalf("expect 110 get %s", r)
	}
}

package leetcode

import "testing"

func TestMultiply(t *testing.T) {
	n1, n2 := "2", "4"
	if r := multiply(n1, n2); r != "8" {
		t.Fatalf("expect 8 get %s", r)
	}

	n1, n2 = "123", "456"
	if r := multiply(n1, n2); r != "56088" {
		t.Fatalf("expect 56088 get %s", r)
	}

	n1, n2 = "1", "99"
	if r := multiply(n1, n2); r != "99" {
		t.Fatalf("expect 99 get %s", r)
	}

	n1, n2 = "22", "33"
	if r := multiply(n1, n2); r != "726" {
		t.Fatalf("expect 726 get %s", r)
	}
	n1, n2 = "9", "9"
	if r := multiply(n1, n2); r != "81" {
		t.Fatalf("expect 81 get %s", r)
	}

	n1, n2 = "9", "99"
	if r := multiply(n1, n2); r != "891" {
		t.Fatalf("expect 891 get %s", r)
	}

	n1, n2 = "0", "999"
	if r := multiply(n1, n2); r != "0" {
		t.Fatalf("expect 0 get %s", r)
	}
}

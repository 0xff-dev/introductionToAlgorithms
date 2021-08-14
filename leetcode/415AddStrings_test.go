package leetcode

import "testing"

func TestAddStrings(t *testing.T) {
	n1, n2 := "11", "123"
	if r := addStrings(n1, n2); r != "134" {
		t.Logf("----%s", r)
		t.Fatalf("expect 134 get %s", r)
	}

	n1, n2 = "456", "77"
	if r := addStrings(n1, n2); r != "533" {
		t.Fatalf("expect 533 get %s", r)
	}

	n1, n2 = "0", "0"
	if r := addStrings(n1, n2); r != "0" {
		t.Fatalf("expect 0 get %s", r)
	}
}

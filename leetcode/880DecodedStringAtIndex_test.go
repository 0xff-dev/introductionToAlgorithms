package leetcode

import "testing"

func TestDecodeAtIndex(t *testing.T) {
	s := "leet2code3"
	k := 10
	if r := decodeAtIndex(s, k); r != "o" {
		t.Fatalf("expect o get %s", r)
	}
	s = "ha22"
	k = 5
	if r := decodeAtIndex(s, k); r != "h" {
		t.Fatalf("expect h get %s", r)
	}
	s = "a2345678999999999999999"
	k = 1
	if r := decodeAtIndex(s, k); r != "a" {
		t.Fatalf("expect a get %s", r)
	}
}

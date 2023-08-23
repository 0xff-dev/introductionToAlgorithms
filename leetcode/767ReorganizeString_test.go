package leetcode

import "testing"

func TestReorganizeString(t *testing.T) {
	if r := reorganizeString("aab"); r != "aba" {
		t.Fatalf("expect 'aba' get %s", r)
	}
	if r := reorganizeString("aaab"); r != "" {
		t.Fatalf("expect '' get %s", r)
	}
	if r := reorganizeString("bbbbb"); r != "" {
		t.Fatalf("expect '' get %s", r)
	}
	if r := reorganizeString("a"); r != "a" {
		t.Fatalf("expect 'a' get %s", r)
	}
}

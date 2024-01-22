package leetcode

import "testing"

func TestFindSubstringInWraproundString(t *testing.T) {
	s := "a"
	if r := findSubstringInWraproundString(s); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	s = "zab"
	if r := findSubstringInWraproundString(s); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	s = "abcdefghijklmnopqrstuvwxyz"
	if r := findSubstringInWraproundString(s); r != 351 {
		t.Fatalf("expect 377 get %d", r)
	}
	s = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyza"
	if r := findSubstringInWraproundString(s); r != 1053 {
		t.Fatalf("expect 1053 get %d", r)
	}

}

package leetcode

import "testing"

func TestMaxVowels(t *testing.T) {
	s := "ibpbhixfiouhdljnjfflpapptrxgcomvnb"
	exp := 7
	if r := maxVowels(s, 33); r != exp {
		t.Fatalf("expect 7 get %d", r)
	}
	s = "abciiidef"
	exp = 3
	if r := maxVowels(s, 3); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	s = "aeiou"
	exp = 2
	if r := maxVowels(s, 2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := maxVowels(s, 5); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}

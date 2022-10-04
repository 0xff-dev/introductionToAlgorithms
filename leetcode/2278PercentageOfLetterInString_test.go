package leetcode

import "testing"

func TestPercentageLetter(t *testing.T) {
	s, letter := "foobar", 'o'
	if r := percentageLetter(s, byte(letter)); r != 33 {
		t.Fatalf("expect 33 get %d", r)
	}

	s, letter = "jjjj", 'k'
	if r := percentageLetter(s, byte(letter)); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

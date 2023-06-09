package leetcode

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')
	if r := nextGreatestLetter(letters, target); r != 'c' {
		t.Fatalf("expect 'c' get '%c'", r)
	}

	letters = []byte{'c', 'f', 'j'}
	target = 'c'
	if r := nextGreatestLetter(letters, target); r != 'f' {
		t.Fatalf("expect 'f' get '%c'", r)
	}
}

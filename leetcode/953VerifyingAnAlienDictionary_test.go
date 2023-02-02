package leetcode

import "testing"

func TestIsAlienSorted(t *testing.T) {
	words, order := []string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"
	if !isAlienSorted(words, order) {
		t.Fatalf("expect true get false")
	}

	words, order = []string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"
	if isAlienSorted(words, order) {
		t.Fatalf("expect false get true")
	}

	words, order = []string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"
	if isAlienSorted(words, order) {
		t.Fatalf("expect false get true")
	}
}

package leetcode

import "testing"

func TestUniqueMorseRepresentations(t *testing.T) {
	words := []string{"gin", "zen", "gig", "msg"}
	if r := uniqueMorseRepresentations(words); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	words = []string{"a"}
	if r := uniqueMorseRepresentations(words); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

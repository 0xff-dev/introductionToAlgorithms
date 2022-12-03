package leetcode

import "testing"

func TestFrequencySort(t *testing.T) {
	s := "tree"
	if r := frequencySort(s); r != "eert" {
		t.Fatalf("expect 'eert' get %s", r)
	}

	s = "cccaaa"
	if r := frequencySort(s); r != "aaaccc" {
		t.Fatalf("expect 'aaaccc' get %s", r)
	}

	s = "Aabb"
	if r := frequencySort(s); r != "bbAa" {
		t.Fatalf("expect 'bbAa' get %s", s)
	}
}

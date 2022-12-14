package leetcode

import "testing"

func TestMergeAlternately(t *testing.T) {
	w1, w2 := "abc", "pqr"
	if r := mergeAlternately(w1, w2); r != "apbqcr" {
		t.Fatalf("expect 'apbqcr' get %s", r)
	}

	w1, w2 = "ab", "pqrs"
	if r := mergeAlternately(w1, w2); r != "apbqrs" {
		t.Fatalf("expect 'apbprs' get %s", r)
	}

	w1, w2 = "abcd", "pq"
	if r := mergeAlternately(w1, w2); r != "apbqcd" {
		t.Fatalf("expect 'apbqcd' get %s", r)
	}
}

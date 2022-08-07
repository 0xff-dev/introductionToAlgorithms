package leetcode

import "testing"

func TestCountVowelPermutation(t *testing.T) {
	if r := countVowelPermutation(1); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	if r := countVowelPermutation(2); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	if r := countVowelPermutation(5); r != 68 {
		t.Fatalf("expect 68 get %d", r)
	}
}

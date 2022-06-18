package leetcode

import "testing"

func TestCountWords(t *testing.T) {
	words1 := []string{"leetcode", "is", "amazing", "as", "is"}
	words2 := []string{"amazing", "leetcode", "is"}
	if r := countWords(words1, words2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

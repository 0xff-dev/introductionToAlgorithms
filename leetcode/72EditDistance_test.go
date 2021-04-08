package leetcode

import "testing"

func TestMinDistance(t *testing.T) {
	word1, word2 := "horse", "ros"
	if _min := minDistance(word1, word2); _min != 3 {
		t.Fatalf("expect 3 get %d", _min)
	}

	word1, word2 = "intention", "execution"
	if _min := minDistance(word1, word2); _min != 5 {
		t.Fatalf("expect 5 get %d", _min)
	}

	word1, word2 = "park", "spake"
	if _min := minDistance(word1, word2); _min != 3 {
		t.Fatalf("expect 3 get %d", _min)
	}
}

package leetcode

import "testing"

func TestMaxScoreWords(t *testing.T) {
	words := []string{"dog", "cat", "dad", "good"}
	lettres := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if r := maxScoreWords(words, lettres, score); r != 23 {
		t.Fatalf("expect 23 get %d", r)
	}
}

package leetcode

import "testing"

func TestMaxScore(t *testing.T) {
	cards := []int{1, 2, 3, 4, 5, 6, 1}
	if r := maxScore(cards, 3); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}

	cards = []int{2, 2, 2}
	if r := maxScore(cards, 2); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	cards = []int{2, 2, 2}
	if r := maxScore(cards, 3); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	cards = []int{9, 7, 7, 9, 7, 7, 9}
	if r := maxScore(cards, 7); r != 55 {
		t.Fatalf("expect 55 get %d", r)
	}
}

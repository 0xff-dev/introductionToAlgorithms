package leetcode

import "testing"

func TestWinnerOfGame(t *testing.T) {
	s := "AAABABB"
	if !winnerOfGame(s) {
		t.Fatalf("with input AAABABB expect true get false")
	}
	s = "AA"
	if winnerOfGame(s) {
		t.Fatalf("with input AA expect false get true")
	}
	s = "ABBBBBBBAAA"

	if winnerOfGame(s) {
		t.Fatalf("with input ABBBBBBBAAA expect false get true")
	}
}

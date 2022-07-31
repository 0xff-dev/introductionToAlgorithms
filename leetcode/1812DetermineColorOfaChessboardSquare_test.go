package leetcode

import "testing"

func TestSquareIsWhite(t *testing.T) {
	input := "a1"
	if squareIsWhite(input) {
		t.Fatalf("expect false get true")
	}

	input = "h3"
	if !squareIsWhite(input) {
		t.Fatalf("expect true get false")
	}

	input = "c7"
	if squareIsWhite(input) {
		t.Fatalf("expect false get true")
	}
}

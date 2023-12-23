package leetcode

import "testing"

func TestIsPathCrossing(t *testing.T) {
	s := "NES"
	if isPathCrossing(s) {
		t.Fatalf("expect false get true")
	}
	s = "NESWW"
	if !isPathCrossing(s) {
		t.Fatalf("expect true get false")
	}
}

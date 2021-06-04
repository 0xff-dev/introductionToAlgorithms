package leetcode

import "testing"

func TestCanConstruct(t *testing.T) {
	a, b := "aa", "b"
	if canConstruct(a, b) {
		t.Fatalf("expect false get true")
	}
}

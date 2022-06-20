package leetcode

import "testing"

func TestMinimumLengthEncoding(t *testing.T) {
	words := []string{"time", "me", "bell"}
	if r := minimumLengthEncoding(words); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}

	if r := minimumLengthEncoding([]string{"t"}); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	if r := minimumLengthEncoding([]string{"me", "time"}); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	words = []string{"p", "grah", "pwosp"}
	if r := minimumLengthEncoding(words); r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}
	// pwosp#grah#
}

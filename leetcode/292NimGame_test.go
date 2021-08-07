package leetcode

import "testing"

func TestCanWinNim(t *testing.T) {
	n := 4
	if r := canWinNim(n); r {
		t.Fatalf("expect false get true")
	}

	n = 5
	if r := canWinNim(n); !r {
		t.Fatalf("expect true get false")
	}
}

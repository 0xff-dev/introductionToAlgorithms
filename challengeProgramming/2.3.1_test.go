package challengeProgramming

import "testing"

func TestBackpack(t *testing.T) {
	weights, values := []int{2, 1, 3, 2}, []int{2, 3, 4, 2}
	if res := Backpack(weights, values, 5); res != 7 {
		t.Fatalf("expect 7 get %d", res)
	}
}

type subsequence struct {
	S1, S2 string
	Ans    int
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, ts := range []subsequence{
		{"abcd", "becd", 3},
		{"aaaaa", "aaaa", 4},
		{"abcdef", "a", 1},
		{"aaaa", "", 0},
		{"abcd", "tyui", 0},
		{"BDCABA", "ABCBDAB", 4},
	} {
		if res := LongestCommonSubsequence(ts.S1, ts.S2); res != ts.Ans {
			t.Fatalf("s1: %s, s2: %s =====> expcet %d get %d", ts.S1, ts.S2, ts.Ans, res)
		}
	}
}

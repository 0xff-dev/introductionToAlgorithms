package leetcode

import "testing"

func TestMaxConsecutiveAnswers(t *testing.T) {
	a := "TTFF"
	k := 2
	if r := maxConsecutiveAnswers(a, k); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	a = "TFFT"
	k = 1
	if r := maxConsecutiveAnswers(a, k); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	a = "TTFTTFTT"
	k = 1
	if r := maxConsecutiveAnswers(a, k); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}

package leetcode

import "testing"

func TestDistinceNames(t *testing.T) {
	ideas := []string{"coffee", "donuts", "time", "toffee"}
	if r := distinctNames(ideas); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	ideas = []string{"lack", "back"}
	if r := distinctNames(ideas); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

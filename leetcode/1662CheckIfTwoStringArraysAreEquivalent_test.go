package leetcode

import "testing"

func TestArrayStringAreEqual(t *testing.T) {
	w1, w2 := []string{"ab", "c"}, []string{"a", "bc"}
	if !arrayStringsAreEqual(w1, w2) {
		t.Fatalf("expect true get false")
	}

	w1, w2 = []string{"a", "cb"}, []string{"ab", "c"}
	if arrayStringsAreEqual(w1, w2) {
		t.Fatalf("expect false get true")
	}
}

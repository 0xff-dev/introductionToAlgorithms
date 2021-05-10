package leetcode

import "testing"

func TestNumDistinct(t *testing.T) {
	s, t1 := "rabbbit", "rabbit"
	if r := numDistinct(s, t1); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	s, t1 = "babgbag", "bag"
	if r := numDistinct(s, t1); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	s, t1 = "", ""
	if r := numDistinct(s, t1); r != 1 {
		t.Fatalf("expcet 1 get %d", r)
	}
}

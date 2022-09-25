package leetcode

import "testing"

func TestCountAsterisks(t *testing.T) {
	s := "l|*e*et|c**o|*de"
	if r := countAsterisks(s); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	s = "iamprogrammer"
	if r := countAsterisks(s); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	s = "yo|uar|e**|b|e***au|tifull"
	if r := countAsterisks(s); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}

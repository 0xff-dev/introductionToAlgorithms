package leetcode

import "testing"

func TestMinFlips(t *testing.T) {
	a, b, c := 2, 6, 5
	if r := minFlips(a, b, c); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	a, b, c = 4, 2, 7
	if r := minFlips(a, b, c); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	a, b, c = 1, 2, 3
	if r := minFlips(a, b, c); r != 0 {
		t.Fatalf("expect 0  get %d", r)
	}

	a, b, c = 8, 3, 5
	if r := minFlips(a, b, c); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

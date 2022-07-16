package leetcode

import "testing"

func TestNumPrimeArrangements(t *testing.T) {
	if r := numPrimeArrangements(5); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}

	if r := numPrimeArrangements(100); r != 682289015 {
		t.Fatalf("expect 682289015 get %d", r)
	}

	if r := numPrimeArrangements(2); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := numPrimeArrangements(1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	if r := numPrimeArrangements(3); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}

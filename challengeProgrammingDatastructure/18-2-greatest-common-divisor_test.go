package challengeprogrammingdatastructure

import "testing"

func TestGCD(t *testing.T) {
	x, y := 6, 9
	if r := GCD(x, y); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	x, y = 2, 7
	if r := GCD(x, y); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	x, y = 147, 105
	if r := GCD(x, y); r != 21 {
		t.Fatalf("expect 21 get %d", r)
	}
}

package challengeprogrammingdatastructure

import "testing"

func TestModPower(t *testing.T) {
	x, n := 2, 3
	if r := modPow(x, n); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	if r := modPowerFor(x, n); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	x, n = 5, 8
	if r := modPow(x, n); r != 390625 {
		t.Fatalf("expect 390625 get %d", r)
	}
	if r := modPowerFor(x, n); r != 390625 {
		t.Fatalf("expect 390625 get %d", r)
	}
}

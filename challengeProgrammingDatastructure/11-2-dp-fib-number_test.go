package challengeprogrammingdatastructure

import "testing"

func TestFib(t *testing.T) {
	n := 5
	// 1, 1, 2, 3, 5, 8
	if r := Fib(n); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	if r := FibInMem(n); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}

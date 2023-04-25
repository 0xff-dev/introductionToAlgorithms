package challengeprogrammingdatastructure

import "testing"

func TestQueens(t *testing.T) {
	if r := Queens(4); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := Queens(5); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	if r := Queens(6); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	if r := Queens(7); r != 40 {
		t.Fatalf("expect 40 get %d", r)
	}
	if r := Queens(8); r != 92 {
		t.Fatalf("expect 92 get %d", r)
	}
}

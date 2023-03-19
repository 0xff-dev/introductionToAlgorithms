package challengeprogrammingdatastructure

import "testing"

func TestReversePolish(t *testing.T) {
	expr := []string{"1", "2", "+", "3", "4", "-", "*"}
	if r := ReversePolish(expr); r != -3 {
		t.Fatalf("expect -3 get %d", r)
	}

	expr = []string{"1", "2", "+"}
	if r := ReversePolish(expr); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}

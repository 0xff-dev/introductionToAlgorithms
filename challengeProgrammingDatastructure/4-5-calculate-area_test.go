package challengeprogrammingdatastructure

import "testing"

func TestCalculateArea(t *testing.T) {
	s := "\\\\\\/_/\\\\///"
	if r := CalculateArea(s); r != 19 {
		t.Fatalf("expect 19 get %d", r)
	}

	s = "\\_\\__/\\_//"
	if r := CalculateArea(s); r != 14 {
		t.Fatalf("expect 9 get %d", r)
	}
}

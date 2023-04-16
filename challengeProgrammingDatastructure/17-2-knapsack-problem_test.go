package challengeprogrammingdatastructure

import "testing"

func TestBackpack(t *testing.T) {
	v, w := []int{4, 5, 2, 8}, []int{2, 2, 1, 3}
	if r := Backpack(v, w, 5); r != 13 {
		t.Fatalf("expect 13 get %d", r)
	}
	w, v = []int{1, 2, 3, 4}, []int{2, 4, 4, 5}
	if r := Backpack(v, w, 5); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}

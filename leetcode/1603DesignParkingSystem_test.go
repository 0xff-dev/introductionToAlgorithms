package leetcode

import "testing"

func TestConstructor1604(t *testing.T) {
	c := Constructor1604(1, 1, 0)
	if !c.AddCar(1) {
		t.Fatalf("expect true get false")
	}
	if !c.AddCar(2) {
		t.Fatalf("expect true get false")
	}
	if c.AddCar(3) {
		t.Fatalf("expect false get true")
	}
	if c.AddCar(1) {
		t.Fatalf("expect false get true")
	}
}

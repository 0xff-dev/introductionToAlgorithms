package leetcode

import "testing"

func TestChampagneTower(t *testing.T) {
	p, qr, qg := 1, 1, 1
	if r := champagneTower(p, qr, qg); r != 0 {
		t.Fatalf("expect 0 get %f", r)
	}
	p, qr, qg = 2, 1, 1
	if r := champagneTower(p, qr, qg); r != 0.5 {
		t.Fatalf("expect 0.t get %f", r)
	}
	p, qr, qg = 4, 2, 2
	if r := champagneTower(p, qr, qg); r != 0.25 {
		t.Fatalf("expect 0.25 get %f", r)
	}
	p, qr, qg = 100000009, 33, 17
	if r := champagneTower(p, qr, qg); r != 1.000 {
		t.Fatalf("expect 1.00 get %f", r)
	}
}

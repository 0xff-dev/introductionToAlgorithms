package leetcode

import "testing"

func TestMinSpeedOnTime(t *testing.T) {
	d := []int{1, 3, 2}
	h := float64(6)
	if r := minSpeedOnTime(d, h); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	d = []int{1, 3, 2}
	h = 2.7
	if r := minSpeedOnTime(d, h); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	d = []int{1, 3, 2}
	h = 1.9
	if r := minSpeedOnTime(d, h); r != -1 {
		t.Fatalf("expect -1get %d", r)
	}
	d = []int{1, 1, 100000}
	h = 2.01
	if r := minSpeedOnTime(d, h); r != 10000000 {
		t.Fatalf("expect 10000000 get %d", r)
	}
}

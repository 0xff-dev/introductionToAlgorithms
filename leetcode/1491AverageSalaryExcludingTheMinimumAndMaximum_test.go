package leetcode

import "testing"

func TestAverage(t *testing.T) {
	salay := []int{4000, 3000, 1000, 2000}
	if r := average(salay); r != 2500.00000 {
		t.Fatalf("expect 2500.0000 get %.4f", r)
	}
}

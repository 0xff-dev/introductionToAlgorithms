package leetcode

import "testing"

func TestGarbageCollection(t *testing.T) {
	garbage := []string{
		"G", "P", "GP", "GG",
	}
	travel := []int{2, 4, 3}
	if r := garbageCollection(garbage, travel); r != 21 {
		t.Fatalf("expect 21 get %d", r)
	}

	garbage = []string{
		"MMM", "PGM", "GP",
	}
	travel = []int{3, 10}
	if r := garbageCollection(garbage, travel); r != 37 {
		t.Fatalf("expect 37 get %d", r)
	}
}

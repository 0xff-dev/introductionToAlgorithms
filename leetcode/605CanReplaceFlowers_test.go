package leetcode

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	f := []int{1, 0, 0, 0, 1}
	n := 1
	if !canPlaceFlowers(f, n) {
		t.Fatalf("expect true get false")
	}

	f = []int{1, 0, 0, 0, 1}
	n = 2
	if canPlaceFlowers(f, n) {
		t.Fatalf("expect false get true")
	}
	f = []int{1, 0, 1}
	n = 1
	if canPlaceFlowers(f, n) {
		t.Fatalf("expect false get true")
	}
	f = []int{0}
	n = 1
	if !canPlaceFlowers(f, n) {
		t.Fatalf("expect true get false")
	}
	f = []int{0, 1, 0, 1, 0, 0}
	n = 1
	if !canPlaceFlowers(f, n) {
		t.Fatalf("expect true get false")
	}
	f = []int{0, 0, 0}
	n = 2
	if !canPlaceFlowers(f, n) {
		t.Fatalf("expect true get false")
	}
	f = []int{1, 0, 0, 0, 0}
	n = 2
	if !canPlaceFlowers(f, n) {
		t.Fatalf("expect true get false")
	}
}

package leetcode

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	a := []int{0, 1, 0}
	if r := peakIndexInMountainArray(a); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	a = []int{0, 2, 1, 0}
	if r := peakIndexInMountainArray(a); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	a = []int{0, 10, 5, 2}
	if r := peakIndexInMountainArray(a); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 5, 4, 2, 1}
	if r := peakIndexInMountainArray(a); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}

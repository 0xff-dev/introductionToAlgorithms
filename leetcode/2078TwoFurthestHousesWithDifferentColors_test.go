package leetcode

import "testing"

func TestMaxDistance(t *testing.T) {
	colors := []int{1, 1, 1, 6, 1, 1, 1}
	if r := maxDistance(colors); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	colors = []int{1, 8, 3, 8, 3}
	if r := maxDistance(colors); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	colors = []int{0, 1}
	if r := maxDistance(colors); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	// 10
	colors = []int{6, 6, 6, 6, 6, 6, 6, 6, 6, 19, 19, 6, 6}

	if r := maxDistance(colors); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
}

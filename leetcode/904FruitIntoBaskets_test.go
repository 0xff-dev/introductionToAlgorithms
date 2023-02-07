package leetcode

import (
	"testing"
)

func TestTotalFruit(t *testing.T) {
	f := []int{1, 2, 1}
	if r := totalFruit(f); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	f = []int{0, 1, 2, 2}
	if r := totalFruit(f); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	f = []int{1, 2, 3, 2, 2}
	if r := totalFruit(f); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}

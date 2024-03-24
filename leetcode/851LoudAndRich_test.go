package leetcode

import (
	"reflect"
	"testing"
)

func TestLoudAndRich(t *testing.T) {
	richer := [][]int{
		{1, 0}, {2, 1}, {3, 1},
		{3, 7}, {4, 3}, {5, 3}, {6, 3},
	}
	q := []int{3, 2, 5, 4, 6, 1, 7, 0}
	exp := []int{5, 5, 2, 5, 4, 5, 6, 7}
	if r := loudAndRich(richer, q); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	richer = [][]int{}
	q = []int{0}
	exp = []int{0}
	if r := loudAndRich(richer, q); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}

	expect := []int{2, 0, 3}
	if r := kWeakestRows(mat, 3); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}

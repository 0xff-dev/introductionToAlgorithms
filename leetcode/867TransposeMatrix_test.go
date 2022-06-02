package leetcode

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	matrix := [][]int{
		{2, 4, -1},
		{-10, 5, 11},
		{18, -7, 6},
	}
	expect := [][]int{
		{2, -10, 18},
		{4, 5, -7},
		{-1, 11, 6},
	}

	if r := transpose(matrix); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	matrix = [][]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	expect = [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	if r := transpose(matrix); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	matrix = [][]int{
		{1},
	}
	expect = [][]int{
		{1},
	}
	if r := transpose(matrix); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %d", expect, r)
	}

	matrix = [][]int{
		{1, 2, 3},
	}
	expect = [][]int{
		{1},
		{2},
		{3},
	}
	if r := transpose(matrix); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	if r := transpose(expect); !reflect.DeepEqual(matrix, r) {
		t.Fatalf("expect %v get %v", matrix, r)
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	exp := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	if r := findDiagonalOrder(mat); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	mat = [][]int{
		{1, 2},
		{3, 4},
	}
	exp = []int{1, 2, 3, 4}

	if r := findDiagonalOrder(mat); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	mat = [][]int{{1}}
	exp = []int{1}
	if r := findDiagonalOrder(mat); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	mat = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	exp = []int{1, 2, 5, 6, 3, 4, 7, 8}
	if r := findDiagonalOrder(mat); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	mat = [][]int{
		{1},
		{2},
		{3},
		{4},
	}

	exp = []int{1, 2, 3, 4}
	if r := findDiagonalOrder(mat); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	mat = [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}

	exp = []int{1, 2, 3, 5, 4, 6, 7, 8}
	if r := findDiagonalOrder(mat); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	mat = [][]int{
		{1, 2, 3, 4},
	}

	exp = []int{1, 2, 3, 4}
	if r := findDiagonalOrder(mat); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

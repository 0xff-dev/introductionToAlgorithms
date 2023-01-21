package leetcode

import (
	"reflect"
	"testing"
)

func TestOnesMinusZeros(t *testing.T) {
	grid := [][]int{
		{0, 1, 1},
		{1, 0, 1},
		{0, 0, 1},
	}
	exp := [][]int{
		{0, 0, 4},
		{0, 0, 4},
		{-2, -2, 2},
	}
	if r := onesMinusZeros(grid); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	grid = [][]int{
		{1, 1, 1},
		{1, 1, 1},
	}

	exp = [][]int{
		{5, 5, 5},
		{5, 5, 5},
	}
	if r := onesMinusZeros(grid); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

}

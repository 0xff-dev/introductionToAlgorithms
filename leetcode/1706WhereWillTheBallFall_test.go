package leetcode

import (
	"reflect"
	"testing"
)

func TestFindBall(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, -1, -1},
		{1, 1, 1, -1, -1},
		{-1, -1, -1, 1, 1},
		{1, 1, 1, 1, -1},
		{-1, -1, -1, -1, -1},
	}

	exp := []int{1, -1, -1, -1, -1}
	if r := findBall(grid); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	grid = [][]int{
		{-1},
	}
	exp = []int{-1}
	if r := findBall(grid); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	grid = [][]int{
		{1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1},
		{1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1},
	}

	exp = []int{0, 1, 2, 3, 4, -1}
	if r := findBall(grid); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

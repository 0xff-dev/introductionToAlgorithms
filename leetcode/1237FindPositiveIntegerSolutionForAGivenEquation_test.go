package leetcode

import (
	"reflect"
	"testing"
)

func f1(a, b int) int {
	return a + b
}
func f2(a, b int) int {
	return a * b
}

func f3(a, b int) int {
	return a - b
}

func TestFindSolution(t *testing.T) {
	exp := [][]int{{4, 1}, {3, 2}, {2, 3}, {1, 4}}
	if r := findSolution(f1, 5); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	exp = [][]int{
		{480, 1}, {240, 2}, {160, 3}, {120, 4}, {96, 5}, {80, 6}, {60, 8}, {48, 10},
		{40, 12}, {32, 15}, {30, 16}, {24, 20}, {20, 24}, {16, 30}, {15, 32},
		{12, 40}, {10, 48}, {8, 60}, {6, 80}, {5, 96}, {4, 120}, {3, 160}, {2, 240}, {1, 480},
	}
	if r := findSolution(f2, 480); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	exp = [][]int{{5, 1}, {1, 5}}
	if r := findSolution(f2, 5); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

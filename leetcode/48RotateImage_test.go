package leetcode

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	r := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	rotate(input)
	for row := 0; row < len(input); row++ {
		t.Logf("%v\n", input[row])
	}
	if !reflect.DeepEqual(input, r) {
		t.Fatal("rotate error")
	}

	input = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	r = [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	rotate(input)
	for row := 0; row < len(input); row++ {
		t.Logf("%v\n", input[row])
	}
	if !reflect.DeepEqual(input, r) {
		t.Fatal("rotate error")
	}

	input = [][]int{
		{1, 2},
		{3, 4},
	}
	r = [][]int{
		{3, 1},
		{4, 2},
	}
	rotate(input)
	for row := 0; row < len(input); row++ {
		t.Logf("%v\n", input[row])
	}
	if !reflect.DeepEqual(input, r) {
		t.Fatal("rotate error")
	}

	input = [][]int{
		{3},
	}
	r = [][]int{
		{3},
	}
	rotate(input)
	for row := 0; row < len(input); row++ {
		t.Logf("%v\n", input[row])
	}
	if !reflect.DeepEqual(input, r) {
		t.Fatal("rotate error")
	}

	input = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	r = [][]int{
		{21, 16, 11, 6, 1},
		{22, 17, 12, 7, 2},
		{23, 18, 13, 8, 3},
		{24, 19, 14, 9, 4},
		{25, 20, 15, 10, 5},
	}
	rotate(input)
	for row := 0; row < len(input); row++ {
		t.Logf("%v\n", input[row])
	}
	if !reflect.DeepEqual(input, r) {
		t.Fatal("rotate error")
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	expect := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}
	if r := floodFill(image, 1, 1, 2); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}

	image = [][]int{
		{1},
	}

	expect = [][]int{
		{2},
	}
	if r := floodFill(image, 0, 0, 2); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}

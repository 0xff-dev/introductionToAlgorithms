package leetcode

import (
	"reflect"
	"testing"
)

func TestFindWinners(t *testing.T) {
	matches := [][]int{
		{1, 3},
		{2, 3},
		{3, 6},
		{5, 6},
		{5, 7},
		{4, 5},
		{4, 8},
		{4, 9},
		{10, 4},
		{10, 9},
	}
	exp := [][]int{
		{1, 2, 10},
		{4, 5, 7, 8},
	}
	if r := findWinners(matches); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	matches = [][]int{
		{2, 3},
		{1, 3},
		{5, 4},
		{6, 4},
	}

	exp = [][]int{
		{1, 2, 5, 6},
		{},
	}
	if r := findWinners(matches); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

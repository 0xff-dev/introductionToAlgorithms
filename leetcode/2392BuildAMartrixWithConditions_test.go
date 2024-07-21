package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildMatrix(t *testing.T) {
	k, rowConditions, colConditions := 3, [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}}
	exp := [][]int{
		{0, 0, 1},
		{3, 0, 0},
		{0, 2, 0},
	}
	if r := buildMatrix(k, rowConditions, colConditions); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	k, rowConditions, colConditions = 3, [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, [][]int{{2, 1}}
	exp = nil
	if r := buildMatrix(k, rowConditions, colConditions); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func TestColorBorder(t *testing.T) {
	grid, row, col, color, exp := [][]int{{1, 1}, {1, 2}}, 0, 0, 3, [][]int{{3, 3}, {3, 2}}
	if r := colorBorder(grid, row, col, color); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	grid, row, col, color, exp = [][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3, [][]int{{1, 3, 3}, {2, 3, 3}}
	if r := colorBorder(grid, row, col, color); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

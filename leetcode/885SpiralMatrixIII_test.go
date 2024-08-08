package leetcode

import (
	"reflect"
	"testing"
)

func TestSpiralMatrixIII(t *testing.T) {
	rows, cols, r, c, exp := 1, 4, 0, 0, [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	if r := spiralMatrixIII(rows, cols, r, c); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	rows, cols, r, c, exp = 5, 6, 1, 4, [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}}
	if r := spiralMatrixIII(rows, cols, r, c); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

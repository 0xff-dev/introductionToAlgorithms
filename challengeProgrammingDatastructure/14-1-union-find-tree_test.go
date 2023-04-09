package challengeprogrammingdatastructure

import (
	"reflect"
	"testing"
)

func TestUnionFindTree(t *testing.T) {
	n, q := 5, 12
	options := [][]int{
		{0, 1, 4},
		{0, 2, 3},
		{1, 1, 2},
		{1, 3, 4},
		{1, 1, 4},
		{1, 3, 2},
		{0, 1, 3},
		{1, 2, 4},
		{1, 3, 0},
		{0, 0, 4},
		{1, 0, 2},
		{1, 3, 0},
	}
	exp := []int{0, 0, 1, 1, 1, 0, 1, 1}
	if r := UnionFindTree(n, q, options); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

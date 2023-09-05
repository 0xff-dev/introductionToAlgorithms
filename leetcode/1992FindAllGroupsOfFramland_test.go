package leetcode

import (
	"reflect"
	"testing"
)

func TestFindFarmland(t *testing.T) {
	land := [][]int{
		{1, 0, 0}, {0, 1, 1}, {0, 1, 1},
	}
	exp := [][]int{
		{0, 0, 0, 0}, {1, 1, 2, 2},
	}
	if r := findFarmland(land); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	land = [][]int{
		{1, 1}, {1, 1},
	}
	exp = [][]int{{0, 0, 1, 1}}
	if r := findFarmland(land); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	land = [][]int{{0}}
	exp = [][]int{}
	if r := findFarmland(land); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

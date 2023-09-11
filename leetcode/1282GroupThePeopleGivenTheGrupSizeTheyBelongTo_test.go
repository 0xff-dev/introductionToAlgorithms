package leetcode

import (
	"reflect"
	"testing"
)

func TestGroupThePeople(t *testing.T) {
	size := []int{3, 3, 3, 3, 3, 1, 3}
	exp := [][]int{
		{0, 1, 2}, {3, 4, 6}, {5},
	}
	if r := groupThePeople(size); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	size = []int{2, 1, 3, 3, 3, 2}
	exp = [][]int{
		{0, 5}, {1}, {2, 3, 4},
	}
	if r := groupThePeople(size); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func TestGetOrder(t *testing.T) {
	tasks := [][]int{
		{1, 2}, {2, 4}, {3, 2}, {4, 1},
	}
	exp := []int{0, 2, 3, 1}
	if r := getOrder(tasks); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tasks = [][]int{
		{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2},
	}
	exp = []int{4, 3, 2, 0, 1}
	if r := getOrder(tasks); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	tasks = [][]int{
		{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24},
	}
	exp = []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}
	if r := getOrder(tasks); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

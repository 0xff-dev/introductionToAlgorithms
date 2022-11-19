package leetcode

import (
	"reflect"
	"testing"
)

func TestLUckyNumbers(t *testing.T) {
	matrix := [][]int{
		{3, 7, 8},
		{9, 11, 13},
		{15, 16, 17},
	}
	exp := []int{15}
	if r := luckyNumbers(matrix); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	matrix = [][]int{
		{1, 10, 4, 2},
		{9, 3, 8, 7},
		{15, 16, 17, 12},
	}

	exp = []int{12}
	if r := luckyNumbers(matrix); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	matrix = [][]int{
		{7, 8},
		{1, 2},
	}

	exp = []int{7}
	if r := luckyNumbers(matrix); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

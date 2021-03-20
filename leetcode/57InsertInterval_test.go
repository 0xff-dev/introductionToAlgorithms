package leetcode

import (
	"reflect"
	"testing"
)

func TestInter(t *testing.T) {
	input := [][]int{
		{1, 3},
		{6, 9},
	}
	interval := []int{2, 5}
	r := insert(input, interval)
	result := [][]int{{1, 5}, {6, 9}}
	if !reflect.DeepEqual(r, result) {
		t.Fatalf("expect %v get %v", result, r)
	}

	input = [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}

	interval = []int{4, 8}
	r = insert(input, interval)
	result = [][]int{{1, 2}, {3, 10}, {12, 16}}
	if !reflect.DeepEqual(r, result) {
		t.Fatalf("expect %v get %v", result, r)
	}

	input = [][]int{}
	interval = []int{5, 7}
	result = [][]int{interval}
	r = insert(input, interval)
	if !reflect.DeepEqual(r, result) {
		t.Fatalf("execpt %v get %v", result, r)
	}

	input = [][]int{{1, 5}}
	interval = []int{2, 3}
	r = insert(input, interval)
	if !reflect.DeepEqual(r, input) {
		t.Fatalf("expect %v get %v", input, r)
	}

	input = [][]int{{1, 5}}
	interval = []int{2, 7}
	r = insert(input, interval)
	result = [][]int{{1, 7}}
	if !reflect.DeepEqual(r, result) {
		t.Fatalf("expect %v get %v", result, r)
	}

	input = [][]int{
		{1, 10},
		{12, 15},
		{16, 18},
	}
	interval = []int{2, 12}
	result = [][]int{{1, 15}, {16, 18}}
	r = insert(input, interval)
	if !reflect.DeepEqual(result, r) {
		t.Fatalf("expect %v get %v", result, r)
	}

	input = [][]int{
		{1, 20},
	}
	interval = []int{3, 8}
	result = input
	r = insert(input, interval)
	if !reflect.DeepEqual(r, result) {
		t.Fatalf("expect %v get %v", result, r)
	}

	input = [][]int{
		{4, 8},
	}
	interval = []int{1, 2}
	result = [][]int{
		{1, 2},
		{4, 8},
	}
	r = insert(input, interval)
	if !reflect.DeepEqual(r, result) {
		t.Fatalf("expect %v get %v", result, r)
	}

	input = [][]int{
		{1, 3},
	}
	interval = []int{4, 8}
	result = [][]int{
		{1, 3},
		{4, 8},
	}
	r = insert(input, interval)
	if !reflect.DeepEqual(r, result) {
		t.Fatalf("expect %v get %v", result, r)
	}
}

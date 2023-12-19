package leetcode

import (
	"reflect"
	"testing"
)

func TestImageSmoother(t *testing.T) {
	img := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	exp := [][]int{
		{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
	}
	if r := imageSmoother(img); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	img = [][]int{
		{100, 200, 100},
		{200, 50, 200},
		{100, 200, 100},
	}
	exp = [][]int{
		{137, 141, 137}, {141, 138, 141}, {137, 141, 137},
	}
	if r := imageSmoother(img); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	img = [][]int{
		{1},
	}
	exp = [][]int{
		{1},
	}
	if r := imageSmoother(img); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

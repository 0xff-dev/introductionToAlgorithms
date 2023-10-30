package leetcode

import (
	"reflect"
	"testing"
)

func TestSortByBits(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	exp := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}
	if r := sortByBits(arr); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	arr = []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	exp = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	if r := sortByBits(arr); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

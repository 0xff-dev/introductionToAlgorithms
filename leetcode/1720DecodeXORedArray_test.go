package leetcode

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	encoded, first := []int{1, 2, 3}, 1
	expect := []int{1, 0, 2, 1}
	if r := decode(encoded, first); !reflect.DeepEqual(r, expect) {
		t.Fatalf("expect %v get %v", expect, r)
	}
}

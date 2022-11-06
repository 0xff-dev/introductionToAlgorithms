package leetcode

import (
	"reflect"
	"testing"
)

func TestDecode1734(t *testing.T) {
	encoded := []int{3, 1}
	exp := []int{1, 2, 3}
	if r := decode1734(encoded); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	encoded = []int{6, 5, 4, 6}
	exp = []int{2, 4, 1, 5, 3}
	if r := decode1734(encoded); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

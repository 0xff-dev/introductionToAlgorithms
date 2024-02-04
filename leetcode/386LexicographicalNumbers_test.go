package leetcode

import (
	"reflect"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	n := 13
	exp := []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}
	if r := lexicalOrder(n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 2
	exp = []int{1, 2}
	if r := lexicalOrder(n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

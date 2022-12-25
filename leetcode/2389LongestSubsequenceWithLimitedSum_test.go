package leetcode

import (
	"reflect"
	"testing"
)

func TestAnswerQueries(t *testing.T) {
	n, q := []int{4, 5, 2, 1}, []int{3, 10, 21}
	exp := []int{2, 3, 4}
	if r := answerQueries(n, q); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n, q = []int{2, 3, 4, 5}, []int{1}
	exp = []int{0}
	if r := answerQueries(n, q); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

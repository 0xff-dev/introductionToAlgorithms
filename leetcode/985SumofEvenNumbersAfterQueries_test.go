package leetcode

import (
	"reflect"
	"testing"
)

func TestSumEvenAfterQueries(t *testing.T) {
	n := []int{1, 2, 3, 4}
	q := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	exp := []int{8, 6, 2, 4}
	if r := sumEvenAfterQueries(n, q); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n = []int{1}
	q = [][]int{{4, 0}}
	exp = []int{0}
	if r := sumEvenAfterQueries(n, q); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

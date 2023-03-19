package challengeprogrammingdatastructure

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	tt := []int{1, 3, 5}
	exp := []int{1, 3, 5}
	if r := BinarySearch(s, tt); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	s = []int{1}
	tt = []int{3, 4, 5, 1}
	exp = []int{1}
	if r := BinarySearch(s, tt); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

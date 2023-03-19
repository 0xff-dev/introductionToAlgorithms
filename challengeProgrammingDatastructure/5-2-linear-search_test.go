package challengeprogrammingdatastructure

import (
	"reflect"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	ss, tt := []int{1, 2, 3, 4, 5}, []int{3, 4, 1}
	exp := []int{1, 3, 4}
	if r := LinearSearch(ss, tt); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

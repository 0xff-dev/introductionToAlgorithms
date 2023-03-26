package challengeprogrammingdatastructure

import (
	"reflect"
	"testing"
)

func TestReconstructTree(t *testing.T) {
	p := []int{1, 2, 3, 4, 5}
	in := []int{3, 2, 4, 1, 5}
	exp := []int{3, 4, 2, 5, 1}
	if r := ReconstructTree(p, in); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	p = []int{1}
	in = []int{1}
	exp = []int{1}
	if r := ReconstructTree(p, in); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

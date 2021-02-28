package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	input := []int{1, 2, 3}
	if r := plusOne(input); !reflect.DeepEqual(r, []int{1, 2, 4}) {
		t.Fatalf("expect [1, 2, 4] get %v", r)
	}

	input = []int{9, 9}
	if r := plusOne(input); !reflect.DeepEqual(r, []int{1, 0, 0}) {
		t.Fatalf("expect [1,0,0] get %v", r)
	}

	input = []int{9}
	if r := plusOne(input); !reflect.DeepEqual(r, []int{1, 0}) {
		t.Fatalf("expect [1,0] get %v", r)
	}
}

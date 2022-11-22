package leetcode

import (
	"reflect"
	"testing"
)

func TestSortPeople(t *testing.T) {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	exp := []string{"Mary", "Emma", "John"}
	if newNames := sortPeople(names, heights); !reflect.DeepEqual(exp, newNames) {
		t.Fatalf("expect %v get %v", exp, newNames)
	}

	names = []string{"Alice", "Bob", "Bob"}
	heights = []int{155, 185, 150}
	exp = []string{"Bob", "Alice", "Bob"}
	if newNames := sortPeople(names, heights); !reflect.DeepEqual(exp, newNames) {
		t.Fatalf("expect %v get %v", exp, newNames)
	}
}

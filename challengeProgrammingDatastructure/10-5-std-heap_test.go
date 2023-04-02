package challengeprogrammingdatastructure

import (
	"reflect"
	"testing"
)

func TestTryPriorityQuque(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exp := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	if r := TryPriorityQuque(nums); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

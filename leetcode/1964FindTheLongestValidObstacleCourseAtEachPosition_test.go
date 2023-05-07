package leetcode

import (
	"reflect"
	"testing"
)

func TestLongestObstacleCourseAtEachPosition(t *testing.T) {
	obstacles := []int{1, 2, 3, 2}
	exp := []int{1, 2, 3, 3}
	if r := longestObstacleCourseAtEachPosition(obstacles); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	obstacles = []int{2, 2, 1}
	exp = []int{1, 2, 1}
	if r := longestObstacleCourseAtEachPosition(obstacles); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	obstacles = []int{3, 1, 5, 6, 4, 2}
	exp = []int{1, 1, 2, 3, 2, 2}
	if r := longestObstacleCourseAtEachPosition(obstacles); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

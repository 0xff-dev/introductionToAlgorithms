package leetcode

import (
	"reflect"
	"testing"
)

func TestSurvivedRobotHealths(t *testing.T) {
	p, h := []int{5, 4, 3, 2, 1}, []int{2, 17, 9, 15, 10}
	d := "RRRRR"
	exp := []int{2, 17, 9, 15, 10}
	if r := survivedRobotsHealths(p, h, d); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	p, h = []int{3, 5, 2, 6}, []int{10, 10, 15, 12}
	d = "RLRL"
	exp = []int{14}
	if r := survivedRobotsHealths(p, h, d); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	p, h = []int{1, 2, 5, 6}, []int{10, 10, 11, 11}
	d = "RLRL"
	exp = []int{}
	if r := survivedRobotsHealths(p, h, d); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

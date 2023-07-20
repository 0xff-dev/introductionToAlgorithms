package leetcode

import (
	"reflect"
	"testing"
)

func TestAsteroidCollistion(t *testing.T) {
	as := []int{5, 10, -5}
	exp := []int{5, 10}
	if r := asteroidCollision(as); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	as = []int{8, -8}
	exp = []int{}
	if r := asteroidCollision(as); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	as = []int{10, 2, -5}
	exp = []int{10}
	if r := asteroidCollision(as); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	as = []int{1, -1, -2, -2}
	exp = []int{-2, -2}
	if r := asteroidCollision(as); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}

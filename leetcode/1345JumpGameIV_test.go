package leetcode

import "testing"

func TestMinJumps(t *testing.T) {
	arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	if r := minJumps(arr); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	arr = []int{7}
	if r := minJumps(arr); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	arr = []int{7, 6, 9, 6, 9, 6, 9, 7}
	if r := minJumps(arr); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}

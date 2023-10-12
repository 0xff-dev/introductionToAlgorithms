package leetcode

import "testing"

func TestFindInMountainArray(t *testing.T) {
	m := MountainArray{data: []int{1, 2, 3, 4, 5, 3, 1}}
	if r := findInMountainArray(3, &m); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	m = MountainArray{data: []int{0, 1, 2, 4, 2, 1}}
	if r := findInMountainArray(3, &m); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
	m = MountainArray{data: []int{1, 2, 3, 4, 5, 3, 1}}
	if r := findInMountainArray(2, &m); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	m = MountainArray{data: []int{1, 5, 2}}
	if r := findInMountainArray(5, &m); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := findInMountainArray(2, &m); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	m = MountainArray{data: []int{1, 2, 3, 5, 3}}

	if r := findInMountainArray(0, &m); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
	if r := findInMountainArray(1, &m); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}

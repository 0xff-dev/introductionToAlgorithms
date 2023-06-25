package leetcode

import "testing"

func TestCountRoutes(t *testing.T) {
	locations, start, finish, fuel := []int{2, 3, 6, 8, 4}, 1, 3, 5
	if r := countRoutes(locations, start, finish, fuel); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	locations, start, finish, fuel = []int{22, 74, 92, 86, 12, 68, 64, 19, 79, 10, 69, 13, 62, 18, 87, 88, 33, 96, 78, 73, 57, 42, 91, 17, 55, 26, 27, 67, 60, 46, 72, 41}, 30, 29, 47
	if r := countRoutes(locations, start, finish, fuel); r != 535415296 {
		t.Fatalf("expect 535415296 get %d", r)
	}
}

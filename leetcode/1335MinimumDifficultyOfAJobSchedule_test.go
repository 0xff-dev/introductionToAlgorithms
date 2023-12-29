package leetcode

import "testing"

func TestMinDifficulty(t *testing.T) {
	j := []int{6, 5, 4, 3, 2, 1}
	d := 2
	if r := minDifficulty(j, d); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	j = []int{9, 9, 9}
	d = 4
	if r := minDifficulty(j, d); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
	j = []int{1, 1, 1}
	d = 3
	if r := minDifficulty(j, d); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	j = []int{100, 23, 48, 96, 72, 33, 67, 300, 235, 786, 999, 1, 0, 3, 5}
	d = 10
	if r := minDifficulty(j, d); r != 1308 {
		t.Fatalf("expect 1308 get %d", r)
	}
	d = 5
	if r := minDifficulty(j, d); r != 1008 {
		t.Fatalf("expect 1008 get %d", r)
	}
}

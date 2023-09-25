package leetcode

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	times := [][]int{
		{2, 1, 1}, {2, 3, 1}, {3, 4, 1},
	}
	n, k := 4, 2
	if r := networkDelayTime(times, n, k); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	times = [][]int{
		{1, 2, 1},
	}
	n, k = 2, 1
	if r := networkDelayTime(times, n, k); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	times = [][]int{
		{1, 2, 1},
	}
	n, k = 2, 2
	if r := networkDelayTime(times, n, k); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}

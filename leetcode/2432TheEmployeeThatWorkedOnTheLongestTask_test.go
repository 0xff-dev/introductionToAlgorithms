package leetcode

import "testing"

func TestHadestWorker(t *testing.T) {
	n := 70
	logs := [][]int{
		{36, 3}, {1, 5}, {12, 8}, {25, 9}, {53, 11}, {29, 12}, {52, 14},
	}
	if r := hardestWorker(n, logs); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}
}

package challengeProgramming

import "testing"

func TestMarkNode(t *testing.T) {
	input := []int{1, 7, 15, 20, 30, 50, 100}
	MarkNode(input, 10)
	input = []int{1, 2, 3, 4, 5}
	MarkNode(input, 3)
	MarkNode(input, 10)
}

func TestTotalCount(t *testing.T) {
	input := []int{8, 5, 8}
	if res := TotalCount(input); res != 34 {
		t.Fatalf("expect 34 get %d", res)
	}
	input = []int{}
	if res := TotalCount(input); res != 0 {
		t.Fatalf("expect 0 get %d", res)
	}
	input = []int{1}
	if res := TotalCount(input); res != 0 {
		t.Fatalf("expect 0 get %d", res)
	}
}

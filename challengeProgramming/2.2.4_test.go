package challengeProgramming

import "testing"

func TestMarkNode(t *testing.T) {
	input := []int{1, 7, 15, 20, 30, 50, 100}
	MarkNode(input, 10)
	input = []int{1, 2, 3, 4, 5}
	MarkNode(input, 3)
	MarkNode(input, 10)
}

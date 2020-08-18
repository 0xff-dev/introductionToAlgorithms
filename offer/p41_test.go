package offer

import (
	"fmt"
	"testing"
)

func TestFindNumberWithSum(t *testing.T) {
	input := []int{1, 2, 4, 7, 11, 15}
	start, end, found := FindNumberWithSum(input, 15)
	if !found {
		t.Fatalf("start: %d, end: %d", 2, 4)
	}
	t.Logf("found start: %d, end: %d", start, end)
}

func TestFindMultiNumberWithSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := FindMultiNumberWithSum(input, 15)
	fmt.Println(result)
}

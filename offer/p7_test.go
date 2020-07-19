package offer

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{5, 3, 9, 6, 1, 2, 3}
	QuickSort(input)
	fmt.Println(input)

	input2 := []int{5, 3, 9, 6, 1, 2, 3}
	QuickSort1(input2)
	fmt.Println(input2)
}

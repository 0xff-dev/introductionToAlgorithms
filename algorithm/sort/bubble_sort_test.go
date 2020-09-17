package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{10, 1, 89, 234, 7423, 87421, 567, 1354678, 8423}
	BubbleSort(input)
	fmt.Println(input)
}

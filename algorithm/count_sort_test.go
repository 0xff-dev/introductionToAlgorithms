package algorith

import (
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {
	input := []int{10, 2, 7, 11, 23, 11, 2, 7}
	res := CountSort(input)
	fmt.Println(res)
}

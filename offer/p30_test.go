package offer

import (
	"fmt"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	data := []int{3, 1, 2, 4, 5}
	fmt.Println(GetLeastNumbers(data, 5))
}

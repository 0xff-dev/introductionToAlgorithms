package dp

import (
	"fmt"
	"testing"
)

func TestMinMatrixCalculateCount(t *testing.T) {
	matrixs := []matrix{
		{
			30,
			35,
		},
		{
			35,
			15,
		},
		{
			15,
			5,
		},
		{
			5,
			10,
		},
		{
			10,
			20,
		},
		{
			20,
			25,
		},
	}
	for k, v := range MinMatrixCalculateCount(matrixs) {
		fmt.Println(fmt.Sprintf("key: %s ---> val: %d", k, v))
	}
}

package offer

import (
	"fmt"
	"testing"
)

func TestExistOneInN(t *testing.T) {
	n := "21345"
	fmt.Println(ExistOneInN(n))
	n = "5"
	fmt.Println(ExistOneInN(n))
}

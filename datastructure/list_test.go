package datastructure

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type testType struct {
	Int int
}

func (self *testType) Less(other interface{}) bool {
	otherObj, _ := other.(*testType)
	return self.Int < otherObj.Int
}

func (self *testType) Swap(other interface{}) {
	otherObj, _ := other.(*testType)
	self.Int, (*otherObj).Int = (*otherObj).Int, self.Int
}
func (self *testType) Print() {
	fmt.Println("test type: ", self.Int)
}

func TestList_Sort(t *testing.T) {
	list := New()
	for cnt := 0; cnt < 6; cnt++ {
		num := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
		list.InsertTail(&testType{num})
	}
	list.Print()
	list.Sort()
	fmt.Println("-------- sort ----------")
	list.Print()
}

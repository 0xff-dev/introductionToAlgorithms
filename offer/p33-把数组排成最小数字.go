package offer

import (
	"fmt"
	"sort"
)

type maxNumArray []int

func (mna maxNumArray) Len() int {
	return len(mna)
}
func (mna maxNumArray) Less(i, j int) bool {
	return fmt.Sprintf("%d%d", mna[i], mna[j]) < fmt.Sprintf("%d%d", mna[j], mna[i])
}
func (mna maxNumArray) Swap(i, j int) {
	mna[i], mna[j] = mna[j], mna[i]
}

func SortMaxNum(nums []int) {
	sort.Sort(maxNumArray(nums))
	fmt.Println(nums)
}

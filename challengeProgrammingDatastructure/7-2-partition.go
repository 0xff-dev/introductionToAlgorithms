package challengeprogrammingdatastructure

import "fmt"

func Partition(nums []int, left, right int) int {
	target := nums[right]
	pre := left - 1

	for index := left; index < right; index++ {
		if nums[index] <= target {
			pre++
			nums[index], nums[pre] = nums[pre], nums[index]
		}
	}
	nums[right], nums[pre+1] = nums[pre+1], nums[right]
	for i := left; i <= right; i++ {
		if i == pre+1 {
			fmt.Printf("[%d] ", nums[i])
			continue
		}
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()
	return pre + 1
}

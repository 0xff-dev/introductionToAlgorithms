package challengeprogrammingdatastructure

func CountingSort(nums []int) []int {
	max := -1
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	target := make([]int, max+1)
	for _, n := range nums {
		target[n]++
	}
	// 统计小于i得数量
	for i := 1; i <= max; i++ {
		target[i] += target[i-1]
	}

	// 0, 0, 1, 1, 2,2
	// 0:2, 1:4, 2:6
	ans := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		ans[target[nums[i]]-1] = nums[i]
		target[nums[i]]--
	}
	return ans
}

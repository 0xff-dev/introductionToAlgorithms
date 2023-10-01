package leetcode

/*
	func countOfOne477(x int) int {
		ans := 0
		for x > 0 {
			x = x & (x - 1)
			ans++
		}
		return ans
	}
*/
func countOfOne477(x int, v *[32]int) {
	index := 31
	for ; x > 0; index-- {
		if x&1 == 1 {
			v[index]++
		}
		x >>= 1
	}
}

func totalHammingDistance(nums []int) int {
	ans := 0
	l := len(nums)
	//cache := make(map[int]int)
	ones := [32]int{}
	for i := 0; i < l; i++ {
		countOfOne477(nums[i], &ones)
	}
	for i := 0; i < 32; i++ {
		if ones[i] == 0 {
			continue
		}

		ans += ones[i] * (l - ones[i])
	}
	/*
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				a := nums[i] ^ nums[j]
				if v, ok := cache[a]; ok {
					ans += v
					continue
				}
				v := countOfOne477(a)
				ans += v
				cache[a] = v
			}
		}
	*/
	return ans
}

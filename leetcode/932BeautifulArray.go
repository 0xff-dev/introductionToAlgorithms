package leetcode

func beautifulArray(n int) []int {
	memo := make(map[int][]int)

	var helper func(int) []int
	helper = func(size int) []int {
		if size == 1 {
			return []int{1}
		}

		if res, ok := memo[size]; ok {
			return res
		}

		res := make([]int, size)
		index := 0

		// 分治：左侧放变换后的奇数，右侧放变换后的偶数
		// 奇数部分长度: (size + 1) / 2
		// 偶数部分长度: size / 2
		for _, x := range helper((size + 1) / 2) {
			res[index] = 2*x - 1
			index++
		}
		for _, x := range helper(size / 2) {
			res[index] = 2 * x
			index++
		}

		memo[size] = res
		return res
	}

	return helper(n)
}

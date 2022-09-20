package leetcode

// 最长公共子串
func findLength(nums1 []int, nums2 []int) int {
	ln1, ln2 := len(nums1), len(nums2)
	ans := 0
	if ln1 == 0 || ln2 == 0 {
		return ans
	}

	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, ln1+1), make([]int, ln1+1)

	/*
			1, 2, 3, 2, 1
		3   0  0  1  1  0
		2   0  1  1  2  0
		4   0  0  0  0  0
		1   1  0  0  0  1
		7

		    0 0 0
		0   1 1 1
		0   1 2 2
		0   1 2 3
	*/

	// 3, 2, 1, 4, 7
	loop := 0
	//tmp := make([][]int, 0)
	for out := 0; out < ln2; out++ {
		for in := 1; in <= ln1; in++ {
			dp[loop][in] = 0
			if nums1[in-1] == nums2[out] {
				dp[loop][in] = dp[1-loop][in-1] + 1
				if dp[loop][in] > ans {
					ans = dp[loop][in]
				}
			}
		}
		//tmp = append(tmp, dp[loop])
		loop = 1 - loop
	}
	/*
		fmt.Println("n1:", nums1)
		for idx, row := range tmp {
			fmt.Println(nums2[idx], row)
		}
	*/
	return ans
}

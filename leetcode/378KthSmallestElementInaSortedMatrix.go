package leetcode

/*
	典型 3 -> 3
	1 3 5
	2 4 6
	7 8 9
*/
func kthSmallest(matrix [][]int, k int) int {
	length := len(matrix)
	if length == 0 {
		return matrix[0][0]
	}
	if length == k*k {
		return matrix[length-1][length-1]
	}
	if length == 1 {
		return matrix[0][k-1]
	}
	// 上半部分或者下半部分

	ans := matrix[0]

	for row := 1; row < length; row++ {
		ans = kthMerge(ans, matrix[row])
	}
	return ans[k-1]
}

func kthMerge(left, right []int) []int {
	ll, lr := len(left), len(right)
	ans := make([]int, ll+lr)
	l, r, idx := 0, 0, -1
	for l < ll || r < lr {
		idx++
		if l >= ll {
			ans[idx] = right[r]
			r++
			continue
		}
		if r >= lr {
			ans[idx] = left[l]
			l++
			continue
		}
		if left[l] < right[r] {
			ans[idx] = left[l]
			l++
			continue
		}
		ans[idx] = right[r]
		r++
	}

	return ans
}

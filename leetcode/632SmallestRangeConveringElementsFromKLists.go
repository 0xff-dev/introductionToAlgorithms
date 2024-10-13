package leetcode

import "math"

/*
func smallestRange(nums [][]int) []int {
	// 我的理解是：对于每个元素都检查一下，然后判断一个范围
	// 我只需要使用第一行就可以，然后一次判断，如果要包含着元素，我需要多大的空间，在其他的行
	l := len(nums)

	first := true
	a, b := 0, 0
	for _, n := range nums[0] {
		ol, or := false, false
		lmin, rmax := n, n
		ta, tb := n, n
		fmt.Printf("calculate n: %d\n", n)
		for line := 1; line < l; line++ {
			ll := len(nums[line])
			idx := sort.Search(ll, func(i int) bool {
				return nums[line][i] >= n
			})
			if idx == ll {
				ol = true
				lmin = min(lmin, nums[line][ll-1])
				continue
			}
			if idx == 0 {
				or = true
				rmax = max(rmax, nums[line][0])
				continue
			}
			lmin = min(lmin, nums[line][idx-1])
			rmax = max(rmax, nums[line][idx])
		}

		if ol && !or {
			ta, tb = lmin, n
		} else if !ol && or {
			ta, tb = n, rmax
		} else if ol && or {
			ta, tb = lmin, rmax
		} else {
			d1 := n - lmin
			d2 := rmax - n
			if d1 == d2 {
				ta, tb = lmin, n
			} else if d1 < d2 {
				ta, tb = lmin, n
			} else {
				ta, tb = n, rmax
			}
			// 计算两套，lmin, n, n, rmax
		}
		fmt.Printf("lmin: %d, rmax: %d, ta: %d, tb=%d,   a: %d, b: %d\n", lmin, rmax, ta, tb, a, b)
		if first {
			// 比较
			a, b = ta, tb
		} else {
			d1 := tb - ta
			d2 := b - a
			if d1 < d2 || (d1 == d2 && ta < a) {
				a, b = ta, tb
			}
		}
		fmt.Printf("got: [%d, %d]\n", a, b)
		first = false
	}
	return []int{a, b}
}
*/

func smallestRange(nums [][]int) []int {
	k := len(nums)
	indices := make([]int, k)
	rangeResult := []int{0, math.MaxInt32}

	for {
		curMin, curMax := math.MaxInt32, math.MinInt32
		minListIndex := 0

		for i := 0; i < k; i++ {
			currentElement := nums[i][indices[i]]

			if currentElement < curMin {
				curMin = currentElement
				minListIndex = i
			}

			if currentElement > curMax {
				curMax = currentElement
			}
		}

		if curMax-curMin < rangeResult[1]-rangeResult[0] {
			rangeResult[0] = curMin
			rangeResult[1] = curMax
		}

		indices[minListIndex]++
		if indices[minListIndex] == len(nums[minListIndex]) {
			break
		}
	}

	return rangeResult
}

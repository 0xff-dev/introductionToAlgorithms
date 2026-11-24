package leetcode

func innerGCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func innerLCM(a, b int64) int64 {
	g := innerGCD(a, b)
	return (a / g) * b
}

func maxLength(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ret, left := 1, 0
	var prod int64 = 1
	var curGCD, curLCM int64 = 0, 0

	for right := 0; right < n; right++ {
		x := int64(nums[right])
		prod *= x
		if right == left {
			curGCD = x
			curLCM = x
		} else {
			curGCD = innerGCD(curGCD, x)
			curLCM = innerLCM(curLCM, x)
		}

		for left <= right && prod > curLCM*curGCD {
			prod /= int64(nums[left])
			left++
			if left <= right {
				curGCD = int64(nums[left])
				curLCM = int64(nums[left])
				prod = int64(nums[left])
				for k := left + 1; k <= right; k++ {
					v := int64(nums[k])
					prod *= v
					curGCD = innerGCD(curGCD, v)
					curLCM = innerLCM(curLCM, v)
				}
			} else {
				prod = 1
				curGCD, curLCM = 0, 0
			}
		}

		if curGCD > 0 && prod == curLCM*curGCD {
			if length := right - left + 1; length > ret {
				ret = length
			}
		}
	}
	return ret
}

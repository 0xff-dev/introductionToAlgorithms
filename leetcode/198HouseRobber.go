package leetcode

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	res := nums[0]
	maxIndex := 0
	for i := 1; i < length; i++ {
		if i != 1 {
			for j := maxIndex; j < i-1; j++ {
				if nums[j] > nums[maxIndex] {
					maxIndex = j
				}
			}
			nums[i] += nums[maxIndex]
			if nums[i] > res {
				res = nums[i]
			}
		} else {
			if nums[i] > res {
				res = nums[i]
			}
		}
	}
	return res
}

func rob2(nums []int) int {
    length := len(nums)
    if length == 0 {
        return 0
    }
    _select, _noSelect := nums[0], 0
    for idx := 1; idx < length; idx++ {
        tmp := _noSelect
        if _select > _noSelect {
            _noSelect = _select
        }
        _select = tmp + nums[idx]
    }
    if _select > _noSelect {
        return _select
    }
    return _noSelect
}

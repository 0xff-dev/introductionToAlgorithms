package leetcode

func deleteAndEarn(nums []int) int {
	r := make([]int, 10001)
	avoid, using, prev := 0, 0, -1
	for _, item := range nums {
		r[item]++
	}

	for i := 0; i <= 10000; i++ {
		if r[i] > 0 {
			_max := using
			if avoid > _max {
				_max = avoid
			}
			if i-1 != prev {
				using = i*r[i] + _max
			} else {
				using = avoid + i*r[i]
			}
			prev, avoid = i, _max
		}
	}
	if avoid > using {
		return avoid
	}
	return using
}

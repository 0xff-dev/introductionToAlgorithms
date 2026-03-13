package leetcode

import (
	"math"
)

/*
const eps3296 = 1e-7

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	mountainHeightInt64 := int64(mountainHeight)
	minTime := int64(math.MaxInt64)
	workerTimesInt64 := make([]int64, len(workerTimes))
	for i := range workerTimes {
		workerTimesInt64[i] = int64(workerTimes[i])
		minTime = min(minTime, workerTimesInt64[i])
	}

	ok := func(t int64) bool {
		// workerTimesInt64[i] * (1 + 2 + 3 + .. n) = t

		reduce := int64(0)
		for i := range workerTimesInt64 {
			reduce += int64((-1.0+math.Sqrt(1+float64(t/workerTimesInt64[i])*8))/2 + eps3296)
		}
		return reduce >= mountainHeightInt64
	}
	l, r := int64(1), mountainHeightInt64/minTime+1
	var ret int64
	for l <= r {
		mid := (l + r) / 2
		if ok(mid) {
			ret = mid
			r = mid - 1
			continue
		}
		l = mid + 1
	}
	return ret
}
*/

const eps = 1e-7

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxWorkerTimes := 0
	for _, t := range workerTimes {
		if t > maxWorkerTimes {
			maxWorkerTimes = t
		}
	}

	l := int64(1)
	r := int64(maxWorkerTimes) * int64(mountainHeight) * int64(mountainHeight+1) / 2
	var ans int64 = 0

	for l <= r {
		mid := (l + r) / 2
		var cnt int64 = 0

		for _, t := range workerTimes {
			work := mid / int64(t)
			// find the largest k such that 1+2+...+k <= work
			k := int64((-1.0+math.Sqrt(1+float64(work)*8))/2 + eps)
			cnt += k
		}
		if cnt >= int64(mountainHeight) {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

package leetcode

func gcd3625(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return gcd3625(b, a%b)
}

func lineK(x1, y1, x2, y2 int) ([2]int, [2]int) {
	if x1 == x2 {
		// 垂直与x轴
		return [2]int{1, 0}, [2]int{x1, 1}
	}

	dx := x2 - x1
	dy := y2 - y1

	// 计算k的分子分母
	g := gcd3625(dy, dx)
	dy /= g
	dx /= g

	// 计算z的分子分母
	// y = dy / dx * x + z
	// z =  (y * dx - dy * x) / dx
	// z = y1 - k * x1 = (y1 * dx - dy * x1) / dx
	zNum := y1*dx - dy*x1
	zDen := dx

	// 化简z
	gz := gcd3625(zNum, zDen)
	zNum /= gz
	zDen /= gz
	return [2]int{dy, dx}, [2]int{zNum, zDen}
}

func countTrapezoids3625(points [][]int) int {
	l := len(points)
	var (
		k, z [2]int
		ret  int
	)
	lines := make(map[[2]int]map[[2]int]int)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			k, z = lineK(points[i][0], points[i][1], points[j][0], points[j][1])
			if _, ok := lines[k]; !ok {
				lines[k] = map[[2]int]int{}
			}
			if _, ok := lines[k][z]; !ok {
				lines[k][z] = 1
			}
			lines[k][z]++
		}
	}

	for _, yOffset := range lines {
		count := make([]int, len(yOffset))
		index, sum := 0, 0
		for _, c := range yOffset {
			count[index] = (c - 1) * c / 2
			sum += count[index]
			index++
		}
		for _, n := range count {
			sum -= n
			ret += sum * n
		}

	}
	return ret
}

func countTrapezoidsII(points [][]int) int {
	n := len(points)
	inf := 1e9 + 7
	slopeToIntercept := make(map[float64][]float64)
	midToSlope := make(map[float64][]float64)
	ans := 0

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx := x1 - x2
			dy := y1 - y2

			var k, b float64
			if x2 == x1 {
				k = inf
				b = float64(x1)
			} else {
				k = float64(y2-y1) / float64(x2-x1)
				b = float64(y1*dx-x1*dy) / float64(dx)
			}

			mid := float64((x1+x2)*10000 + (y1 + y2))
			slopeToIntercept[k] = append(slopeToIntercept[k], b)
			midToSlope[mid] = append(midToSlope[mid], k)
		}
	}

	for _, sti := range slopeToIntercept {
		if len(sti) == 1 {
			continue
		}

		cnt := make(map[float64]int)
		for _, bVal := range sti {
			cnt[bVal]++
		}

		totalSum := 0
		for _, count := range cnt {
			ans += totalSum * count
			totalSum += count
		}
	}

	for _, mts := range midToSlope {
		if len(mts) == 1 {
			continue
		}

		cnt := make(map[float64]int)
		for _, kVal := range mts {
			cnt[kVal]++
		}

		totalSum := 0
		for _, count := range cnt {
			ans -= totalSum * count
			totalSum += count
		}
	}

	return ans
}

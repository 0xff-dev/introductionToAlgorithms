package leetcode

func gcd1232(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd1232(y, x%y)
}

func kd(x1, y1, x2, y2 int) (int, int, int, int) {
	dx := x2 - x1
	dy := y2 - y1
	if dx == 0 {
		return dy, -1, -1, -1
	}
	g1 := gcd1232(dx, dy)
	kx := dx / g1
	ky := dy / g1

	// d = b - a*k k=ky/kx
	dy1 := y1*kx - x1*ky
	dx1 := kx
	g2 := gcd1232(dx1, dx1)

	return ky, kx, dy1 / g2, dx1 / g2
}
func checkStraightLine(coordinates [][]int) bool {
	x := coordinates[0][0] - coordinates[1][0]
	y := coordinates[0][1] - coordinates[1][1]
	if x == 0 {
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][0] != coordinates[0][0] {
				return false
			}
		}
		return true
	}
	if y == 0 {
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][1] != coordinates[0][1] {
				return false
			}
		}
		return true
	}

	a, b, c, d := 0, 0, 0, 0
	for i := 1; i < len(coordinates); i++ {
		a1, b1, c1, d1 := kd(coordinates[i-1][0], coordinates[i-1][1], coordinates[i][0], coordinates[i][1])
		if i == 1 {
			a, b, c, d = a1, b1, c1, d1
			continue
		}
		if a != a1 || b != b1 || c != c1 || d != d1 {
			return false
		}
	}
	return true
}

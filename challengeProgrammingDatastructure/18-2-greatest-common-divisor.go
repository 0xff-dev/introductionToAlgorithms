package challengeprogrammingdatastructure

func GCD(x, y int) int {
	for x%y != 0 {
		x, y = y, x%y
	}
	return y
}

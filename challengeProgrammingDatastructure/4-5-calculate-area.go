package challengeprogrammingdatastructure

func CalculateArea(area string) int {
	ans := 0
	gen := make([]int, 0)
	for i, which := range area {
		if which == '\\' {
			gen = append(gen, i)
			continue
		}
		if which == '/' && len(gen) > 0 {
			left := gen[len(gen)-1]
			gen = gen[:len(gen)-1]
			diff := i - left // little area
			ans += diff
		}
	}
	return ans
}

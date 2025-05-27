package leetcode

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start == destination {
		return 0
	}
	n := len(distance)
	cur, end := start, (start+1)%n
	r := 0
	for cur != destination {
		r += distance[cur]
		cur, end = end, (end+1)%n
	}

	r1 := 0
	cur, end = start, (start-1+n)%n
	for cur != destination {
		r1 += distance[end]
		cur, end = end, (end-1+n)%n
	}
	return min(r, r1)
}

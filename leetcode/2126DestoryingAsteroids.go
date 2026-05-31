package leetcode

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for i := range asteroids {
		if mass < asteroids[i] {
			return false
		}
		mass += asteroids[i]
	}

	return true
}

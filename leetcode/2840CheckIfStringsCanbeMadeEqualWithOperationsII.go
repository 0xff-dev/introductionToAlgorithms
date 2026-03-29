package leetcode

import "sort"

func checkStrings(s1 string, s2 string) bool {
	a1, b1 := make([]byte, 0), make([]byte, 0)
	n := len(s1)
	for i := 0; i < n; i += 2 {
		a1 = append(a1, s1[i])
		b1 = append(b1, s2[i])
	}

	sort.Slice(a1, func(i, j int) bool {
		return a1[i] < a1[j]
	})
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})

	for i := range a1 {
		if a1[i] != b1[i] {
			return false
		}
	}

	a2, b2 := make([]byte, 0), make([]byte, 0)
	for i := 1; i < n; i += 2 {
		a2 = append(a2, s1[i])
		b2 = append(b2, s2[i])
	}
	sort.Slice(a2, func(i, j int) bool {
		return a2[i] < a2[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	for i := range a2 {
		if a2[i] != b2[i] {
			return false
		}
	}
	return true
}

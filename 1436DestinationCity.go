package leetcode

func destCity(paths [][]string) string {
	from := make(map[string]struct{})
	to := make(map[string]struct{})
	for _, path := range paths {
		from[path[0]] = struct{}{}
		to[path[1]] = struct{}{}
	}
	for v := range to {
		if _, ok := from[v]; !ok {
			return v
		}
	}
	return ""
}

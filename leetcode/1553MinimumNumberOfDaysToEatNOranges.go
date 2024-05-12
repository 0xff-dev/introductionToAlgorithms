package leetcode

func minDays(n int) int {
	/*
		cache := map[int]int{
			1: 1, 2: 2, 3: 2,
		}
		var dfs func(int) int
		dfs = func(now int) int {
			if v, ok := cache[now]; ok {
				return v
			}
			cur := -1
			if now%2 == 0 {
				r := dfs(now / 2)

				if cur == -1 || cur > r {
					cur = r
				}
			}
			if now%3 == 0 {
				r := dfs(now/3)
				if cur == -1 || cur > r {
					cur = r
				}
			}
			r := dfs(now - 1)
			if cur == -1 || cur > r {
				cur = r
			}

			cache[now] = cur + 1
			return cur + 1
		}
		return dfs(n)
	*/
	// dfs不行，试试bfs
	in := make(map[int]struct{})
	queue := [][2]int{{n, 0}}
	for len(queue) > 0 {
		nq := make([][2]int, 0)
		for _, cur := range queue {
			if cur[0] == 1 {
				return cur[1] + 1
			}
			next := cur[0] - 1
			if _, ok := in[next]; !ok {
				in[next] = struct{}{}
				nq = append(nq, [2]int{next, cur[1] + 1})
			}
			if cur[0]%2 == 0 {
				next = cur[0] / 2
				if _, ok := in[next]; !ok {
					in[next] = struct{}{}
					nq = append(nq, [2]int{next, cur[1] + 1})
				}
			}

			if cur[0]%3 == 0 {
				next = cur[0] / 3
				if _, ok := in[next]; !ok {
					in[next] = struct{}{}
					nq = append(nq, [2]int{next, cur[1] + 1})
				}
			}
		}
		queue = nq
	}
	return -1
}

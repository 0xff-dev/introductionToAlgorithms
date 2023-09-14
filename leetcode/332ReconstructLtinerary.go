package leetcode

import "sort"

func findItinerary(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] < tickets[j][1]
		}
		return tickets[i][0] < tickets[j][0]
	})

	to := make(map[string][]string)
	for _, t := range tickets {
		a, b := t[0], t[1]
		if _, ok := to[a]; !ok {
			to[a] = make([]string, 0)
		}
		to[a] = append(to[a], b)
	}
	left := len(tickets)
	var (
		ans  []string
		anss string
	)

	var dfs func(string, string, []string, int)
	dfs = func(station, seq string, path []string, l int) {
		if l == 0 {
			if anss == "" || anss > seq {
				anss = seq
				x := make([]string, len(path))
				copy(x, path)
				ans = x
			}
			return
		}
		for i, next := range to[station] {
			if next == "" {
				continue
			}
			to[station][i] = ""
			dfs(next, seq+next, append(path, next), l-1)
			to[station][i] = next
		}
	}
	dfs("JFK", "", []string{"JFK"}, left)
	return ans
}

func olahuilu(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		res []string
	)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

package leetcode

type state752 struct {
	cur [4]byte
	cnt int
}

// "0201","0101","0102","1212","2002"
// 0202
// 0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
func openLock(deadends []string, target string) int {
	dm := make(map[[4]byte]struct{})
	for _, d := range deadends {
		x := [4]byte{}
		for i := 0; i < 4; i++ {
			x[i] = d[i]
		}
		dm[x] = struct{}{}
	}
	var (
		add func(byte) byte
		del func(byte) byte
	)
	add = func(a byte) byte {
		if a == '9' {
			return '0'
		}
		return a + 1
	}
	del = func(a byte) byte {
		if a == '0' {
			return '9'
		}
		return a - 1
	}
	/*
		cache := make(map[[4]byte]int)
		cache[[4]byte{'0', '0', '0', '0'}] = 0
		var (
			dfs func(int, [4]byte) int
			add func(byte) byte
			del func(byte) byte
		)
		add = func(a byte) byte {
			if a == '9' {
				return '0'
			}
			return a + 1
		}
		del = func(a byte) byte {
			if a == '0' {
				return '9'
			}
			return a - 1
		}
		usedState := make(map[[4]byte]struct{})
		// state来自于前一个状态的某个为切换1下。
		// dfs+cache, 草，dfs计算的有问题，切换bfs
		dfs = func(cost int, state [4]byte) int {
			if _, ok := dm[state]; ok {
				return -1
			}
			if state[0] == '0' && state[1] == '0' && state[2] == '0' && state[3] == '0' {
				return cost
			}
			if v, ok := cache[state]; ok {
				return v
			}
			ans := -1
			usedState[state] = struct{}{}
			for i := 0; i < 4; i++ {
				source := state[i]
				next := add(source)
				state[i] = next
				if _, ok := usedState[state]; !ok {
					if r := dfs(cost+1, state); r != -1 && (ans == -1 || ans > r) {
						ans = r
					}
				}

				next = del(source)
				state[i] = next
				if _, ok := usedState[state]; !ok {
					if r := dfs(cost+1, state); r != -1 && (ans == -1 || ans > r) {
						ans = r
					}
				}
				state[i] = source
			}
			delete(usedState, state)
			cache[state] = ans
			return ans
		}
				return dfs(0, state)
	*/

	state := [4]byte{target[0], target[1], target[2], target[3]}
	queue := []state752{{state, 0}}
	used := make(map[[4]byte]struct{})
	used[state] = struct{}{}
	for len(queue) > 0 {
		nq := make([]state752, 0)
		for _, cur := range queue {
			if _, ok := dm[cur.cur]; ok {
				continue
			}
			if cur.cur[0] == '0' && cur.cur[1] == '0' && cur.cur[2] == '0' && cur.cur[3] == '0' {
				return cur.cnt
			}

			for i := 0; i < 4; i++ {
				ss := cur.cur[i]
				source := cur.cur
				source[i] = add(ss)
				if _, ok := used[source]; !ok {
					used[source] = struct{}{}
					nq = append(nq, state752{cur: source, cnt: cur.cnt + 1})
				}
				source[i] = del(ss)
				if _, ok := used[source]; !ok {
					used[source] = struct{}{}
					nq = append(nq, state752{cur: source, cnt: cur.cnt + 1})
				}
			}
		}
		queue = nq
	}
	return -1
}

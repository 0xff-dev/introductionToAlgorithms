package leetcode

type UnionFind527 struct {
	Father []int
}

func (uf *UnionFind527) FindFather(x int) int {
	if uf.Father[x] != x {
		uf.Father[x] = uf.FindFather(uf.Father[x])
	}
	return uf.Father[x]
}

func (uf *UnionFind527) Union(x, y int) {
	xf := uf.Father[x]
	yf := uf.Father[y]
	if xf < yf {
		uf.Father[yf] = xf
		return
	}
	uf.Father[xf] = yf
}

func findCircleNum(isConnected [][]int) int {
	cities := len(isConnected)
	uf := UnionFind527{Father: make([]int, cities+1)}
	for i := 1; i <= cities; i++ {
		// 自己连自己
		uf.Father[i] = i
	}
	for row := 1; row <= cities; row++ {
		for col := 1; col <= cities; col++ {
			if isConnected[row-1][col-1] == 1 {
				f1 := uf.FindFather(row)
				f2 := uf.FindFather(col)
				uf.Union(f1, f2)
			}
		}
	}
	provinces := make(map[int]struct{})
	for i := 1; i <= cities; i++ {
		provinces[uf.FindFather(i)] = struct{}{}
	}
	return len(provinces)
}

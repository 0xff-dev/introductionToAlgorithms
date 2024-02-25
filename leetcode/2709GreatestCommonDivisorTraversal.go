package leetcode

// 靠，dfs绝对超时，如果要所有节点都互联，并查集，集合到一个点
/*
func gcd2709(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
			continue
		}
		b = b - a
	}
	return a
}

*/
type unionFind2709 struct {
	father []int
}

func (u *unionFind2709) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}
func (u *unionFind2709) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

/*
func isPrime2709(num int, gcdCache map[int]bool) bool {
	if v, ok := gcdCache[num]; ok {
		return v
	}
	gcdCache[num] = false
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	gcdCache[num] = true
	return true
}
func canTraverseAllPairs(nums []int) bool {
	l := len(nums)
	gcdCache := make(map[int]bool)
	u := unionFind2709{father: make([]int, l)}
	isPrime := make([]bool, l)
	for i := 0; i < l; i++ {
		isPrime[i] = isPrime2709(nums[i], gcdCache)
		u.father[i] = i
	}

	shouldUnion := make(map[[2]int]struct{})
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if _, ok := shouldUnion[[2]int{nums[i], nums[j]}]; ok {
				u.union(i, j)
				continue
			}
			if (!isPrime[i] && !isPrime[j]) || (nums[i]%nums[j] == 0 || nums[j]%nums[i] == 0) {
				u.union(i, j)
				shouldUnion[[2]int{nums[i], nums[j]}] = struct{}{}
				shouldUnion[[2]int{nums[j], nums[i]}] = struct{}{}
			}
		}
	}
	for i := 0; i < l; i++ {
		if u.find(i) != 0 {
			return false
		}
	}
	return true
}
*/

func canTraverseAllPairs(nums []int) bool {
	l := len(nums)
	if l == 1 {
		return false
	}
	alloc := 100000
	primeFactor := make([]int, alloc+1)
	for i := 2; i <= alloc; i++ {
		if primeFactor[i] == 0 {
			for next := i; next <= alloc; next += i {
				primeFactor[next] = i
			}
		}
	}

	u := unionFind2709{father: make([]int, 2*alloc+1)}
	for i := 0; i < 2*alloc+1; i++ {
		u.father[i] = i
	}
	for _, n := range nums {
		if n == 1 {
			return false
		}
		cur := n
		for cur > 1 {
			fac := primeFactor[cur]
			root := fac + alloc
			u.union(root, n)
			for cur%fac == 0 {
				cur /= fac
			}
		}
	}
	root := u.find(nums[0])
	for i := 1; i < l; i++ {
		if u.find(nums[i]) != root {
			return false
		}
	}
	return true
}

package leetcode

const MX3629 = 1000001

var factors [MX3629][]int

func init3629() {
	for i := 2; i < MX3629; i++ {
		if len(factors[i]) == 0 {
			for j := i; j < MX3629; j += i {
				factors[j] = append(factors[j], i)
			}
		}
	}
}

func minJumps3629(nums []int) int {
	n := len(nums)
	edges := make(map[int][]int)
	for i, a := range nums {
		if len(factors[a]) == 1 {
			edges[a] = append(edges[a], i)
		}
	}
	res := 0
	seen := make([]bool, n)
	seen[n-1] = true
	q := []int{n - 1}
	for {
		var q2 []int
		for _, i := range q {
			if i == 0 {
				return res
			}
			if i > 0 && !seen[i-1] {
				seen[i-1] = true
				q2 = append(q2, i-1)
			}
			if i < n-1 && !seen[i+1] {
				seen[i+1] = true
				q2 = append(q2, i+1)
			}
			for _, p := range factors[nums[i]] {
				if list, ok := edges[p]; ok {
					for _, j := range list {
						if !seen[j] {
							seen[j] = true
							q2 = append(q2, j)
						}
					}
					delete(edges, p)
				}
			}
		}
		q = q2
		res++
	}
}

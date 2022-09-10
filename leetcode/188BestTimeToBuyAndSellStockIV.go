package leetcode

import "container/heap"

type stock []int

func (s *stock) Len() int {
	return len(*s)
}

func (s *stock) Less(i, j int) bool {
	return (*s)[i] > (*s)[j]
}

func (s *stock) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *stock) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *stock) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

type stockPair struct {
	start, end int
}

func maxProfit188(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}

	//diff := 0
	//differences := &stock{}
	//
	//// 3, 2, 3, 6, 5, 0, 3
	//
	//min := prices[0]
	//for idx := 1; idx < len(prices); idx++ {
	//	if prices[idx] < prices[idx-1] {
	//		if diff != 0 && prices[idx] < min {
	//			heap.Push(differences, diff)
	//			diff = 0
	//		}
	//		fmt.Println()
	//		continue
	//	}
	//
	//	diff += prices[idx] - prices[idx-1]
	//	fmt.Printf("diff += prices[%d]-prices[%d]=%d\n", idx, idx-1, diff)
	//}
	//if diff != 0 {
	//	heap.Push(differences, diff)
	//}
	//loop := 0
	//ans := 0
	//for differences.Len() > 0 && loop < k {
	//	pop := heap.Pop(differences).(int)
	//	ans += pop
	//	loop++
	//}

	s, e := 0, -1
	pairs := make([]stockPair, 0)
	profit := &stock{}
	for {
		for s = e + 1; s+1 < len(prices) && prices[s] >= prices[s+1]; s++ {
		}
		for e = s; e+1 < len(prices) && prices[e] <= prices[e+1]; e++ {
		}

		if s == e {
			break
		}

		/*
			情况1：价格[v1]<=价格[v2]，价格[p1]<=价格[p2]。
			在这种情况下，如果我们能进行一次交易，我们将使用（v1，p2）。
			如果我们可以进行两次交易，我们将使用（v1，p1）和（v2，p2）。
			等价地，我们可以把（v1，p2）视为一个交易机会，把（v2，p1）视为另一个交易机会。
			关键的想法是，这两个原始谷峰对提供了两个交易机会：（v1，p2）和（v2，p1）。

			情况2：价格[v1]>=价格[v2]或价格[p1]>=价格[p2]。
			在这种情况下，如果我们可以进行一次交易，我们将使用（v1，p1）或（v2，p2）。
			如果我们可以进行两次交易，我们将同时使用（v1，p1）和（v2，p2）。
			也就是说，这两个谷峰对提供了两个交易机会：（v1，p1）和（v2，p2）。
		*/
		// 2, 3, 1, 5  的情况，合并
		// if v1 <= v2 && p1 <= p2
		for l := len(pairs); l > 0 && prices[s] <= prices[pairs[l-1].start]; l-- {
			heap.Push(profit, prices[pairs[l-1].end]-prices[pairs[l-1].start])
			pairs = pairs[:l-1]
		}

		// if v1 >= v2 || p1 >= p2
		// 1, 3, 2, 5
		for l := len(pairs); l > 0 && prices[e] >= prices[pairs[l-1].end]; l-- {
			heap.Push(profit, prices[pairs[l-1].end]-prices[s])
			s = pairs[l-1].start
			pairs = pairs[:l-1]
		}

		pairs = append(pairs, stockPair{s, e})
	}

	for l := len(pairs); l > 0; l-- {
		heap.Push(profit, prices[pairs[l-1].end]-prices[pairs[l-1].start])
		pairs = pairs[:l-1]
	}

	ans := 0
	loop := 0
	for profit.Len() > 0 && loop < k {
		top := heap.Pop(profit).(int)
		ans += top
		loop++
	}
	return ans
}

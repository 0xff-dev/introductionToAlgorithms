package leetcode

/*
import (
	"container/heap"
	"fmt"
	"sort"
)

type ssItem struct {
	c   byte
	idx int // 在数组原来的位置
	cnt int
}
type ss struct {
	mm   []int
	data []ssItem
}

func (i *ss) Len() int {
	return len(i.data)
}

func (i1 *ss) Swap(i, j int) {
	i1.data[i], i1.data[j] = i1.data[j], i1.data[i]
}
func (i1 *ss) Less(i, j int) bool {
	a := i1.mm[i1.data[i].cnt]
	b := i1.mm[i1.data[j].cnt]
	if a == b {
		return i1.data[i].idx < i1.data[j].idx
	}
	return a < b
}
func (i *ss) Push(x interface{}) {
	i.data = append(i.data, x.(ssItem))
}

func (i *ss) Pop() interface{} {
	old := i.data
	l := len(i.data)
	x := old[l-1]
	i.data = old[:l-1]
	return x
}

func help1531(data *[]ssItem) int {
	sort.Slice(*data, func(i, j int) bool {
		return (*data)[i].idx < (*data)[j].idx
	})

	index := 0
	for idx := 1; idx < len(*data); idx++ {
		if (*data)[idx].c == (*data)[index].c {
			(*data)[index].cnt += (*data)[idx].cnt
			continue
		}
		index++
		(*data)[index] = (*data)[idx]
	}
	*data = (*data)[:index+1]
	cnt := 0
	for _, n := range *data {
		if n.cnt == 1 {
			cnt++
			continue
		}
		if n.cnt < 10 {
			// a1
			cnt += 2
			continue
		}

		if n.cnt < 100 {
			// a11
			cnt += 3
			continue
		}
		cnt += 4
	}
	return cnt
}

// a100这是最大的情况,b
// 可以知道长度为1，2，3位的有多少种按规律删除就可以了。
// 对于a, a2这种，只需要一次就可以减少一位
// a3, a9 则需要 需要 减少n次，才可以减少一位
// a10 一次减少一位
func getLengthOfOptimalCompression(s string, k int) int {
	mm := make([]int, 101)
	mm[1], mm[2], mm[10], mm[100] = 1, 1, 1, 1
	for i := 3; i < 100; i++ {
		if i == 10 {
			continue
		}
		mm[i] = mm[i-1] + 1
	}

	ans := 0
	arr := ss{mm: mm, data: []ssItem{{c: s[0], idx: 0, cnt: 1}}}
	index := 0
	for idx := 1; idx < len(s); idx++ {
		if s[idx] == arr.data[index].c {
			arr.data[index].cnt++
			continue
		}
		index++
		arr.data = append(arr.data, ssItem{c: s[idx], idx: index, cnt: 1})
	}

	fmt.Printf("data: %v\n", arr.data)

	for k > 0 {
		ans = help1531(&arr.data)
		fmt.Printf("----- requeue: %v\n", arr.data)
		heap.Init(&arr)
		ok := true
		for arr.Len() > 0 && k > 0 {
			fmt.Printf("cur data: %v\n", arr.data)
			top := heap.Pop(&arr).(ssItem)
			nk := mm[top.cnt]
			fmt.Printf("top--- cnt: %d, c: %c, nk: %d\n", top.cnt, top.c, nk)

			if k < nk {
				ok = false
				break
			}
			k -= nk
			if nk <= 2 {
				ans -= nk
				break
			}
			top.cnt -= nk
			heap.Push(&arr, top)
		}
		fmt.Println()
		if !ok {
			break
		}
	}
	return ans
}
*/

func min8(a, b int8) int8 {
	if a <= b {
		return a
	}
	return b
}

func getLengthOfOptimalCompression(s string, k int) int {
	n := int8(len(s))
	dp := make([][]int8, n+1)
	dp[0] = make([]int8, k+1)
	for i := int8(0); i < n; i++ {
		dp[i+1] = make([]int8, k+1)
		dp[i+1][0] = n
		copy(dp[i+1][1:], dp[i][0:])
		for j, jEnd := int8(0), int8(k); j <= jEnd; j++ {
			var lastCharCount, differentCharsRemovedCount int8
			for l := i; l >= 0; l-- {
				if s[l] == s[i] {
					lastCharCount++
				} else {
					differentCharsRemovedCount++
					if differentCharsRemovedCount > j {
						break
					}
				}
				var compressedCharsLen int8
				if lastCharCount == 1 {
					compressedCharsLen = 1
				} else if lastCharCount < 10 {
					compressedCharsLen = 2
				} else if lastCharCount < 100 {
					compressedCharsLen = 3
				} else {
					compressedCharsLen = 4
				}
				dp[i+1][j] = min8(dp[i+1][j], dp[l][j-differentCharsRemovedCount]+compressedCharsLen)
			}
		}
	}
	return int(dp[n][k])
}

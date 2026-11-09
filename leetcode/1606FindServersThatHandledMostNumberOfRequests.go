package leetcode

import (
	"container/heap"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type BusyServer struct {
	releaseTime int
	serverId    int
}

type MinHeap1606 []BusyServer

func (h MinHeap1606) Len() int           { return len(h) }
func (h MinHeap1606) Less(i, j int) bool { return h[i].releaseTime < h[j].releaseTime }
func (h MinHeap1606) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap1606) Push(x interface{}) {
	*h = append(*h, x.(BusyServer))
}
func (h *MinHeap1606) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func busiestServers(k int, arrival []int, load []int) []int {
	// 可用服务器集合，用红黑树存储，支持 O(log k) 的查找和删除
	available := redblacktree.NewWithIntComparator()
	for i := 0; i < k; i++ {
		available.Put(i, true)
	}

	busyHeap := &MinHeap1606{}
	heap.Init(busyHeap)

	requestCount := make([]int, k)
	maxRequests := 0

	for i := 0; i < len(arrival); i++ {
		currTime := arrival[i]
		duration := load[i]

		// 步骤 1：释放当前时间点之前已经空闲的服务器
		for busyHeap.Len() > 0 && (*busyHeap)[0].releaseTime <= currTime {
			item := heap.Pop(busyHeap).(BusyServer)
			available.Put(item.serverId, true)
		}

		// 步骤 2：如果没有可用服务器，当前请求被丢弃
		if available.Size() == 0 {
			continue
		}

		// 步骤 3：查找目标服务器 >= (i % k)
		target := i % k
		var serverId int

		// 在红黑树中寻找第一个 >= target 的 key
		node, found := available.Ceiling(target)
		if found {
			serverId = node.Key.(int)
		} else {
			// 如果没有找到 >= target 的，回绕到树中的最小值（即最小的 ID）
			serverId = available.Left().Key.(int)
		}

		// 步骤 4：分配服务器
		available.Remove(serverId)
		heap.Push(busyHeap, BusyServer{
			releaseTime: currTime + duration,
			serverId:    serverId,
		})

		requestCount[serverId]++
		if requestCount[serverId] > maxRequests {
			maxRequests = requestCount[serverId]
		}
	}

	// 步骤 5：收集所有达到最大处理请求数的服务器 ID
	var result []int
	for i := 0; i < k; i++ {
		if requestCount[i] == maxRequests {
			result = append(result, i)
		}
	}

	return result
}

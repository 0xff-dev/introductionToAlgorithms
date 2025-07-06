package leetcode

type FindSumPairs struct {
	c1, c2 map[int]int
	n2     []int
}

func Constructor1865(nums1 []int, nums2 []int) FindSumPairs {
	f := FindSumPairs{
		c1: map[int]int{}, c2: map[int]int{}, n2: nums2,
	}
	for _, n := range nums1 {
		f.c1[n]++
	}
	for _, n := range nums2 {
		f.c2[n]++
	}
	return f
}

func (this *FindSumPairs) Add(index int, val int) {
	source := this.n2[index]
	this.n2[index] += val
	this.c2[source]--
	if this.c2[source] == 0 {
		delete(this.c2, source)
	}
	this.c2[this.n2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	ans := 0
	for k, c := range this.c1 {
		if c1, ok := this.c2[tot-k]; ok {
			ans += c * c1
		}
	}
	return ans
}

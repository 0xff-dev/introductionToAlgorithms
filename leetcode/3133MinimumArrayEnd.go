package leetcode

func minEnd(n int, x int) int64 {
	// x 自己的0位置，保证了下限，你怎么与我AND，这个位的值一直都是0
	// 所吸需要做的就是，给0的位置填1，看几个够就ok了。
	// 最后的AND=x，说明选择的数字是1的位置必须与x相同
	// 需要做的就是，将x的0的位置填充0或者1，够n-1个就可以了。x是第一个要选择的元素
	ans := int64(x)
	if n == 1 {
		return ans
	}
	nn := int64(n) - 1
	zero := 0
	c := int64(0)
	for {
		zero++
		c = 1 << (zero - 1)
		if nn <= c {
			break
		}
		nn -= c
	}
	c |= (nn - 1)
	cur := int64(1)
	for ; c > 0; c >>= 1 {
		for ans&cur != 0 {
			cur <<= 1
		}
		now := c & 1
		if now == 1 {
			ans |= cur
		} else {
			cur <<= 1
		}
	}
	return ans
}

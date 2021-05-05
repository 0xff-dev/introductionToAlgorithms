package leetcode

// 从数组中删除一元素，然后在数组中找到最大和的子自数组
// dp

func maximumSum(arr []int) int {
	length := len(arr)
	if length == 1 {
		return arr[0]
	}

	maxResult := -10001
	c1, c2 := arr[0], arr[0]
	for idx := 1; idx < length; idx++ {
		tmp := c1
		if c1+arr[idx] > arr[idx] {
			c1 += arr[idx]
		} else {
			c1 = arr[idx]
		}
		if c1 > maxResult {
			maxResult = c1
		}

		t := c2
		c2 = tmp
		if t+arr[idx] > c2 {
			c2 = t + arr[idx]
		}
		if c2 > maxResult {
			maxResult = c2
		}

	}
	return maxResult
}

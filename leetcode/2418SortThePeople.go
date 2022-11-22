package leetcode

func sortPeople(names []string, heights []int) []string {
	// quick sort
	quickSort(names, heights, 0, len(heights)-1)
	return names
}

func quickSort(names []string, heights []int, start, end int) {
	if start < end {
		compare := heights[start]
		pre := start
		for walker := start + 1; walker <= end; walker++ {
			if heights[walker] > compare {
				pre++
				heights[walker], heights[pre] = heights[pre], heights[walker]
				names[walker], names[pre] = names[pre], names[walker]
			}
		}
		heights[start], heights[pre] = heights[pre], heights[start]
		names[start], names[pre] = names[pre], names[start]

		quickSort(names, heights, start, pre-1)
		quickSort(names, heights, pre+1, end)
	}
}

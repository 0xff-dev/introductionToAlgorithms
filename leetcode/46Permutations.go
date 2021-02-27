package leetcode

import (
	"fmt"
	"strings"
)

func permute(nums []int) [][]int {
	r := make([][]int, 0)
	check := make(map[string]struct{})
	template := strings.Repeat("%d", len(nums))
	permuteAux(nums, 0, &r, check, template)
	return r
}

func permuteAux(nums []int, index int, r *[][]int, exist map[string]struct{}, template string) {
	if index == len(nums) {
		dest := make([]int, 0)
		interfaceDest := make([]interface{}, 0)
		for _, n := range nums {
			dest = append(dest, n)
			interfaceDest = append(interfaceDest, n)
		}
		key := fmt.Sprintf(template, interfaceDest...)
		if _, ok := exist[key]; ok {
			return
		}
		exist[key] = struct{}{}
		*r = append(*r, dest)
		return
	}
	for idx := index; idx < len(nums); idx++ {
		nums[idx], nums[index] = nums[index], nums[idx]
		permuteAux(nums, index+1, r, exist, template)
		nums[idx], nums[index] = nums[index], nums[idx]
	}
}

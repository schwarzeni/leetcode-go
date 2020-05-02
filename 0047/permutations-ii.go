package _0047

import "sort"

// 2020.05.02
// https://leetcode-cn.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	var f func(level int, idxMap []int)
	var result [][]int

	sort.Ints(nums)

	f = func(level int, idxMap []int) {
		if level == len(nums) {
			r := make([]int, len(nums))
			for i, v := range idxMap {
				r[v-1] = nums[i]
			}
			result = append(result, r)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && idxMap[i-1] == 0 {
				continue
			}
			if idxMap[i] == 0 {
				idxMap[i] = level + 1
				f(level+1, idxMap)
				idxMap[i] = 0
			}
		}
	}

	f(0, make([]int, len(nums)))

	return result
}

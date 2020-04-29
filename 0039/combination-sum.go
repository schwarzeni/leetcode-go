package _0039

import "sort"

// 2020.04.29
// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return sum(candidates, target, []int{})
}

func sum(candidates []int, target int, pre []int) (result [][]int) {
	for idx, c := range candidates {
		if c > target {
			return
		}
		r := make([]int, len(pre)+1)
		copy(r, pre)
		r[len(pre)] = c
		if c == target {
			result = append(result, r)
			return
		} else if c < target {
			for _, rr := range sum(candidates[idx:], target-c, r) {
				result = append(result, rr)
			}
		}
	}
	return
}

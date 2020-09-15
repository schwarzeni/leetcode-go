package _0228

import "strconv"

// 2020.09.15
// https://leetcode-cn.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var res []string
	l := nums[0]
	for i := 0; i < len(nums); i++ {
		if (i < len(nums)-1 && nums[i]+1 != nums[i+1]) || i == len(nums)-1 {
			if l == nums[i] {
				res = append(res, strconv.Itoa(l))
			} else {
				res = append(res, strconv.Itoa(l)+"->"+strconv.Itoa(nums[i]))
			}
			if i < len(nums)-1 {
				l = nums[i+1]
			}
		}
	}
	return res
}

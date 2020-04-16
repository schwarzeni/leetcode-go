package _0015

import (
	"fmt"
	"sort"
)

// 2020.04.16
// https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	var result [][]int
	var record = make(map[string]struct{})
	sort.Ints(nums)

	for ia, a := range nums {
		if a <= 0 {
			ib := ia + 1
			ic := len(nums) - 1
			for ib < ic {
				r := nums[ib] + nums[ic]
				if r < -a {
					ib++
				} else if r > -a {
					ic--
					if ic < 0 {
						break
					}
				} else {
					key := fmt.Sprintf("%d%d%d", a, nums[ib], nums[ic])
					if _, ok := record[key]; !ok {
						result = append(result, []int{a, nums[ib], nums[ic]})
						record[key] = struct{}{}
					}
					ib++
					ic--
				}
			}
		}
	}

	return result
}

package _0018

import "sort"

// 2020.04.22
// https://leetcode-cn.com/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	var (
		i, j, k, l, t int
		result        [][]int
	)

	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)

	for i = 0; i < len(nums)-3; i++ {
		for j = i + 1; j < len(nums)-2; j++ {
			k, l = j+1, len(nums)-1
			t = target - nums[i] - nums[j]
			for k < l {
				if nums[k]+nums[l] == t {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
				} else if nums[k]+nums[l] < t {
					k++
				} else {
					l--
				}
				for k < l && k > j+1 && nums[k] == nums[k-1] {
					k++
				}
				for k < l && l < len(nums)-1 && nums[l+1] == nums[l] {
					l--
				}
			}
			for j < len(nums)-3 && nums[j] == nums[j+1] {
				j++
			}
		}
		for i < len(nums)-4 && nums[i] == nums[i+1] {
			i++
		}
	}

	return result
}

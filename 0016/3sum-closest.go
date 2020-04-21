package _0016

import "sort"

// 2020.04.21
// https://leetcode-cn.com/problems/3sum-closest/
//func threeSumClosest(nums []int, target int) int {
//	sort.Ints(nums)
//	min := target - (nums[0] + nums[1] + nums[2])
//	for i := 0; i < len(nums)-2; i++ {
//		if i < len(nums)-3 && nums[i] == nums[i+1] {
//			continue
//		}
//		for j := i + 1; j < len(nums)-1; j++ {
//			for k := i + 2; k < len(nums); k++ {
//				m := target - nums[i] - nums[j] - nums[k]
//				if m == 0 {
//					return target
//				}
//				if abs(min) > abs(m) {
//					min = m
//				}
//			}
//		}
//	}
//	return target - min
//}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := target - (nums[0] + nums[1] + nums[2])
	for i := 0; i < len(nums)-2; i++ {
		if i < len(nums)-3 && nums[i] == nums[i+1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			threesum := nums[i] + nums[j] + nums[k]
			if abs(target-threesum) < abs(min) {
				min = target - threesum
			}
			if target == threesum {
				return target
			}
			if target < threesum {
				k--
			} else {
				j++
			}
		}
	}
	return target - min
}

func abs(d int) int {
	if d < 0 {
		return -d
	}
	return d
}

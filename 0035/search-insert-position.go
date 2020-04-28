package _0035

// 2020.04.28
// https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1

	if len(nums) == 0 {
		return 0
	}

	for {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			if i >= j {
				return i + 1
			}
			i = mid + 1
		} else {
			if i >= j {
				return i
			}
			j = mid - 1
		}
	}
}

package _0081

// 2020.06.10
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		for l < r && nums[l] == nums[l+1] {
			l++
		}
		for l < r && nums[r] == nums[r-1] {
			r--
		}
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] >= nums[l] {
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}

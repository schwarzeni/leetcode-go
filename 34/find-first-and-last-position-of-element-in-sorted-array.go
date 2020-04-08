package _34

// 2020.04.08
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	var (
		headIdx = -1
		tailIdx = -1
		numSize = len(nums)
		x, y    int
	)

	// find headIdx
	x = 0
	y = numSize - 1
	for {
		if x > y {
			break
		}
		mid := (x + y) / 2
		if nums[mid] < target {
			x = mid + 1
		} else if nums[mid] > target {
			y = mid - 1
		} else if nums[x] <= target {
			if nums[x] == target {
				headIdx = x
				break
			} else if mid-x == 1 {
				headIdx = mid
				break
			} else {
				y = mid
			}
		} else {
			break
		}
	}

	// find tailIdx
	x = 0
	y = numSize - 1
	for {
		if x > y {
			break
		}
		mid := (x + y) / 2
		if nums[mid] < target {
			x = mid + 1
		} else if nums[mid] > target {
			y = mid - 1
		} else if nums[y] >= target {
			if nums[y] == target {
				tailIdx = y
				break
			} else if y-mid == 1 {
				tailIdx = mid
				break
			} else {
				x = mid
			}
		} else {
			break
		}
	}
	return []int{headIdx, tailIdx}
}

// 1 2 3 3 3 4 5 6 7 8
// 0 1 2 3 4 5 6 7 8 9

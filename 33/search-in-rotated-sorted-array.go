package _33

// 2020.04.06
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	var (
		lidx = 0
		ridx = len(nums) - 1
	)
	for {
		if (lidx == ridx && nums[lidx] != target) || lidx > ridx {
			return -1
		}
		if lidx == ridx && nums[lidx] == target {
			return lidx
		}

		var (
			mid = (lidx + ridx) / 2
			//ml  = mid - 1
			//mr  = mid + 1
		)
		//if ml < lidx {
		//	ml = lidx
		//}
		//if mr >= ridx {
		//	mr = ridx
		//}
		if target == nums[mid] {
			return mid
		}
		//if nums[ml] <= nums[mid] && nums[mr] >= nums[mid] {
		if nums[lidx] <= nums[ridx] { // 子数组有序
			if target > nums[mid] {
				lidx = mid + 1
				//continue
			} else {
				ridx = mid - 1
				//continue
			}
		} else { // 子数组无序
			if nums[mid] >= nums[lidx] { // fix test_3
				if nums[lidx] <= target && nums[mid] > target {
					ridx = mid - 1
					//continue
				} else {
					lidx = mid + 1
					//continue
				}
			} else { // fix test_4
				if nums[ridx] >= target && nums[mid] < target {
					lidx = mid + 1
					//continue
				} else {
					ridx = mid - 1
					//continue
				}
			}
		}
		//}
		// 最大的点
		//if nums[ml] <= nums[mid] && nums[mr] <= nums[mid] {
		//	if target > nums[mid] {
		//		return -1
		//	}
		//	if target >= nums[lidx] {
		//		ridx = mid
		//		continue
		//	} else {
		//		lidx = mid + 1
		//		continue
		//	}
		//}
		//// 最小的点
		//if nums[ml] >= nums[mid] && nums[mr] >= nums[mid] {
		//	if target < nums[mid] {
		//		return -1
		//	}
		//	if target >= nums[lidx] {
		//		ridx = mid - 1
		//		continue
		//	} else {
		//		lidx = mid
		//		continue
		//	}
		//}
	}
}

// 跳转点的情况：
// 1. nums[k-1] <= nums[k]
//    nums[k+1] <= nums[k]

// 2. nums[k-1] >= nums[k]
//    nums[k+1] >= nums[k]

// 3. nums[k-1] <= nums[k] <= nums[k+1]

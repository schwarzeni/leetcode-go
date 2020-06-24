package _0162

// 2020.06.24
// https://leetcode-cn.com/problems/find-peak-element/
func findPeakElement(nums []int) int {
	return fpe(0, len(nums)-1, nums)
}

func fpe(l, r int, nums []int) int {
	mid := l + (r-l)/2
	if l == r {
		return mid
	}
	if mid == 0 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}
	lv, midv, rv := nums[mid-1], nums[mid], nums[mid+1]
	if lv < midv && rv < midv {
		return mid
	}
	if (lv < midv && midv < rv) || lv > midv && rv > midv {
		return fpe(mid+1, r, nums)
	}
	if lv > midv && midv > rv {
		if l == mid {
			return mid
		}
		return fpe(l, mid-1, nums)
	}
	return -1
}

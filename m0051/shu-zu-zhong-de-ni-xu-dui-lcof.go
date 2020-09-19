package _m0051

// 2020.09.19
// 逆序对，挺经典的一道题目，使用到了归并排序法
// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(nums []int) int {
	count := 0
	buf := make([]int, len(nums))
	var merge func(l, r int)
	merge = func(l, r int) {
		if r-l+1 <= 2 {
			if r-l+1 == 2 && nums[l] > nums[r] {
				count++
				nums[r], nums[l] = nums[l], nums[r]
			}
			return
		}
		mid := (l + r) / 2
		merge(l, mid-1)
		merge(mid, r)
		lidx, ridx, ll, rl := l, mid, mid, r+1
		for idx := l; lidx < ll || ridx < rl; idx++ {
			if ridx >= rl || (lidx < ll && nums[lidx] <= nums[ridx]) {
				buf[idx] = nums[lidx]
				lidx++
			} else {
				count += ll - lidx
				buf[idx] = nums[ridx]
				ridx++
			}
		}
		copy(nums[l:r+1], buf[l:r+1])
	}
	merge(0, len(nums)-1)
	return count
}

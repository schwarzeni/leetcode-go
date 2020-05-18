package _0152

import "math"

// 2020.05.18
// https://leetcode-cn.com/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	max, imax, imin := math.MinInt64, 1, 1
	for _, v := range nums {
		// 如果数组的数是负数，那么会导致最大的变最小的，最小的变最大的。因此交换两个的值。
		if v < 0 {
			imax, imin = imin, imax
		}
		imax = (int)(math.Max(float64(v), float64(v*imax)))
		imin = (int)(math.Min(float64(v), float64(v*imin)))
		max = (int)(math.Max(float64(max), float64(imax)))
	}
	return max
}

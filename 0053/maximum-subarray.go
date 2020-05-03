package _0053

// 2020.05.03
// https://leetcode-cn.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	cache := make([]int, len(nums))
	max := 0
	for i, v := range nums {
		if i == 0 || cache[i-1] < 0 {
			cache[i] = v
		} else {
			cache[i] = cache[i-1] + v
		}
		if i == 0 || max < cache[i] {
			max = cache[i]
		}
	}
	return max
}

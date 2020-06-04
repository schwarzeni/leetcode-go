package _0238

// 2020.06.04
// https://leetcode-cn.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] = result[i+1] * nums[i]
	}
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] * nums[i]
	}
	for i := 0; i < len(nums); i++ {
		a, b := 1, 1
		if i > 0 {
			a = nums[i-1]
		}
		if i < len(nums)-1 {
			b = result[i+1]
		}
		result[i] = a * b
	}
	return result
}

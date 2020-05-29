package _198

// 2020.05.29
// https://leetcode-cn.com/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	if len(nums) == 1 {
		return nums[0]
	}
	cache := make([]int, len(nums))
	cache[0], cache[1] = nums[0], nums[1]
	if nums[0] > nums[1] {
		cache[1] = cache[0]
	}
	for i := 2; i < len(nums); i++ {
		cache[i] = cache[i-1]
		if nums[i]+cache[i-2] > cache[i-1] {
			cache[i] = nums[i] + cache[i-2]
		}
	}
	return cache[len(nums)-1]
}

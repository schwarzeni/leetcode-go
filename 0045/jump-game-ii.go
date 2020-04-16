package _45

// 2020.04.15
// https://leetcode-cn.com/problems/jump-game-ii/
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cache := make([]int, len(nums))
	cache[0] = 0

	for idx, v := range nums {
		for i := idx + 1; i <= idx+v && i < len(nums); i++ {
			if cache[i] == 0 || cache[i] > cache[idx]+1 {
				cache[i] = cache[idx] + 1
			}
		}
	}

	return cache[len(nums)-1]
}

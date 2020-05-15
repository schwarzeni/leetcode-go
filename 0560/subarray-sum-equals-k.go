package _0560

// 2020.05.15
// https://leetcode-cn.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	cache := make(map[int]int)
	cache[0] = 1
	for _, v := range nums {
		sum += v
		if v, ok := cache[sum-k]; ok {
			count += v
		}
		cache[sum] = cache[sum] + 1
	}
	return count
}

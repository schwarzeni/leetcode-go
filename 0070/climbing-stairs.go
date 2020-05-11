package _0070

// 2020.05.12
// https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	cache := make([]int, n+1)
	cache[0], cache[1] = 1, 1
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}

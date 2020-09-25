package _0204

// 2020.09.25
// https://leetcode-cn.com/problems/count-primes/
func countPrimes(n int) int {
	mark := make([]bool, n+1)
	count := 0
	for i := 2; i < n; i++ {
		if !mark[i] {
			count++
			for j := 2; i*j < n; j++ {
				mark[i*j] = true
			}
		}
	}
	return count
}

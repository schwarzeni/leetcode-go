package _0069

// 2020.05.09
// https://leetcode-cn.com/problems/sqrtx/
func mySqrt(x int) int {
	n := 1
	for {
		if (n-1)*(n-1) <= x && x < n*n {
			return n - 1
		}
		n += 1
	}
}

package _0172

// 2020.09.25
// https://leetcode-cn.com/problems/factorial-trailing-zeroes/
func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		count += n / 5
		n /= 5
	}
	return count
}

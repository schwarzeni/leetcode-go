package _0009

// 2020.04.19
// https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var (
		xx = x
		y  = 0
	)
	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	return y == xx
}

package _0191

// 2020.08.20
// https://leetcode-cn.com/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 > 0 {
			count++
		}
		num >>= 1
	}
	return count
}

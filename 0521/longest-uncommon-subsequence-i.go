package _0521

import "math"

// 2020.05.21
// https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return int(math.Max(float64(len(a)), float64(len(b))))
}

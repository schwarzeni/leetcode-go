package _0171

// 2020.08.21
// https://leetcode-cn.com/problems/excel-sheet-column-number/
func titleToNumber(s string) int {
	res := 0
	for len(s) > 0 {
		res = res*26 + int(s[0]-'A') + 1
		s = s[1:]
	}
	return res
}

package _0013

// 2020.04.20
// https://leetcode-cn.com/problems/roman-to-integer/
var table = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	idx := len(s) - 1
	result := 0

	for idx >= 0 {
		if idx > 0 {
			if s[idx] == 'V' && s[idx-1] == 'I' {
				result += 4
				idx--
			} else if s[idx] == 'X' && s[idx-1] == 'I' {
				result += 9
				idx--
			} else if s[idx] == 'L' && s[idx-1] == 'X' {
				result += 40
				idx--
			} else if s[idx] == 'C' && s[idx-1] == 'X' {
				result += 90
				idx--
			} else if s[idx] == 'D' && s[idx-1] == 'C' {
				result += 400
				idx--
			} else if s[idx] == 'M' && s[idx-1] == 'C' {
				result += 900
				idx--
			} else {
				result += table[s[idx]]
			}
		} else {
			result += table[s[idx]]
		}
		idx--
	}

	return result
}

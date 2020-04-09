package _5

// 2020.04.09
// https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	var (
		strLen      = len(s)
		idx         = 0
		maxStartIdx = 0
		maxEndIdx   = 0
	)
	if strLen == 0 {
		return ""
	}

	for idx <= (strLen-1)*2 {
		var i, j int
		i = idx / 2
		if idx%2 == 0 {
			j = i
		} else {
			j = i + 1
		}
		for {
			if s[i] != s[j] {
				if maxEndIdx-maxStartIdx < j-i-2 {
					maxStartIdx = i + 1
					maxEndIdx = j - 1
				}
				break
			}
			i -= 1
			j += 1
			if i < 0 && j >= strLen {
				return s
			}
			if i < 0 {
				if maxEndIdx-maxStartIdx < j {
					maxStartIdx = 0
					maxEndIdx = j - 1
				}
				break
			}
			if j >= strLen {
				if maxEndIdx-maxStartIdx < strLen-1-i {
					maxStartIdx = i + 1
					maxEndIdx = strLen - 1
				}
				break
			}
		}
		idx += 1
	}
	return s[maxStartIdx : maxEndIdx+1]
}

// 0   1   2   3   4
// b   a   b   a   d
// 0 1 2 3 4 5 6 7 8

// 0   1   2   3
// b   a   b   a
// 0 1 2 3 4 5 6

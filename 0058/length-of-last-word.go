package _0058

// 2020.05.04
// https://leetcode-cn.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	count, size := 0, len(s)-1

	for idx := size; idx >= 0; idx-- {
		if s[idx] != ' ' {
			count++
		} else if count != 0 {
			break
		}
	}

	return count
}

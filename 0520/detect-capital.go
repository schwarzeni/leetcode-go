package _0520

import (
	"unicode"
)

// 2020.05.21
// https://leetcode-cn.com/problems/detect-capital/
func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	status := -1
	if unicode.IsUpper(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
		status = 0 // USA
	}
	if unicode.IsLower(rune(word[0])) && unicode.IsLower(rune(word[1])) {
		status = 1 // leetcode
	}
	if unicode.IsUpper(rune(word[0])) && unicode.IsLower(rune(word[1])) {
		status = 2 // Google
	}
	if status == -1 {
		return false
	}
	for _, s := range word[2:] {
		if status == 0 && unicode.IsLower(s) {
			return false
		}
		if status == 1 && unicode.IsUpper(s) {
			return false
		}
		if status == 2 && unicode.IsUpper(s) {
			return false
		}
	}
	return true
}

package _0151

import "strings"

// 2020.06.10
// https://leetcode-cn.com/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	var wordList []string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			j := i
			for ; j >= 0 && s[j] != ' '; j-- {
			}
			wordList = append(wordList, s[j+1:i+1])
			i = j
		}
	}
	return strings.Join(wordList, " ")
}

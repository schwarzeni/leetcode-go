package _0394

import (
	"bytes"
	"strings"
)

// 2020.05.28
// https://leetcode-cn.com/problems/decode-string/
func decodeString(s string) string {
	_, rs := repeatStr(0, s)
	return rs
}

// 解析 ab23[abc ... edf]
func repeatStr(idx int, s string) (nextIdx int, rs string) {
	count, buffer := 0, bytes.Buffer{}
	for idx < len(s) {
		for idx < len(s) && isLetter(s[idx]) {
			buffer.WriteByte(s[idx])
			idx++
		}
		if idx < len(s) && s[idx] == ']' {
			return idx + 1, buffer.String()
		}
		for idx < len(s) && isDigit(s[idx]) {
			count = count*10 + int(s[idx]-'0')
			idx++
		}
		idx++ // '['
		subrs := ""
		idx, subrs = repeatStr(idx, s)
		buffer.WriteString(strings.Repeat(subrs, count))
		count = 0
	}
	return idx, buffer.String()
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

//func isDigit(c rune) bool {
//	return c >= '0' && c <= '9'
//}
//
//func isLetter(c rune) bool {
//	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
//}

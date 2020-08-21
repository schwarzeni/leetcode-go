package _0168

import "bytes"

// 2020.08.21
// https://leetcode-cn.com/problems/excel-sheet-column-title/
// 可以使用逆向思维，逆推关系式，参考 171
func convertToTitle(n int) string {
	var res bytes.Buffer
	for n > 0 {
		n--
		res.WriteByte(byte(n%26 + int('A')))
		n /= 26
	}
	return reverse(res.String())
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// 27 = 26^1 * 1 + 26^0 * 1
// 28 = 26^1 * 1 + 26^0 * 2
//    = 10^1 * 2 + 10^0 * 8

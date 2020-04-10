package _8

import "math"

// 2020.04.10
// https://leetcode-cn.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
	const (
		S_BEFORE_NUM = iota
		S_IN_NUM
		S_AFTER_NUM
	)
	var (
		result      int
		sLen        = len(str)
		isDigit     = func(data byte) bool { return data >= '0' && data <= '9' }
		numStartIdx = -1
		numEndIdx   = -1
		isNegative  bool
		idx         int
		status      = S_BEFORE_NUM
	)

	for idx < sLen {
		switch status {
		case S_BEFORE_NUM:
			switch {
			case str[idx] == ' ':
				status = S_BEFORE_NUM
				idx += 1
			case str[idx] == '-': // fix test_2
				nextIdx := idx + 1
				if nextIdx < sLen && isDigit(str[nextIdx]) {
					isNegative = true
					status = S_IN_NUM
					idx += 1
					numStartIdx = idx
				} else {
					status = S_AFTER_NUM
				}
			case str[idx] == '+': // fix test_1
				nextIdx := idx + 1
				if nextIdx < sLen && isDigit(str[nextIdx]) {
					isNegative = false
					status = S_IN_NUM
					idx += 1
					numStartIdx = idx
				} else {
					status = S_AFTER_NUM
				}

			case isDigit(str[idx]):
				numStartIdx = idx
				status = S_IN_NUM
			default:
				status = S_AFTER_NUM
			}
		case S_IN_NUM:
			if nextIdx := idx + 1; nextIdx < sLen && isDigit(str[nextIdx]) {
				status = S_IN_NUM
				idx += 1
			} else {
				numEndIdx = idx
				status = S_AFTER_NUM
			}
		case S_AFTER_NUM:
			if numStartIdx == -1 {
				return 0
			}
			for _, ch := range str[numStartIdx : numEndIdx+1] {
				ch -= '0'
				// fix test_3 test9 test10
				if !isNegative && (result > 214748364 || (result == 214748364 && ch > 7)) {
					return math.MaxInt32
				}
				if isNegative && (result > 214748364 || (result == 214748364 && ch > 8)) {
					return math.MinInt32
				}
				result = result*10 + int(ch)
			}
			if isNegative {
				result = -result
			}
			return result
		}
	}

	return result
}

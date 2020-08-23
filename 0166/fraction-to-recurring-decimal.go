package _0166

import "strconv"

// 2020.08.23
// https://leetcode-cn.com/problems/fraction-to-recurring-decimal/
func fractionToDecimal(numerator int, denominator int) string {
	ans := ""
	if numerator == 0 || denominator == 0 {
		return "0"
	}
	// 处理整数部分
	ans += (strconv.Itoa(numerator / denominator))
	if ans == "0" && ((numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0)) {
		ans = "-0"
	}
	numerator = numerator % denominator
	if numerator == 0 {
		return ans
	}
	ans += "."
	if numerator < 0 {
		numerator *= -1
	}
	if denominator < 0 {
		denominator *= -1
	}
	prev := make(map[int]int)
	for {
		if idx, ok := prev[numerator]; ok {
			repeat := ans[idx:]
			ans = ans[:idx]
			ans += "(" + repeat + ")"
			break
		}
		res := (numerator * 10) / denominator
		remain := (numerator * 10) % denominator
		if remain == 0 {
			ans += strconv.Itoa(res)
			break
		}
		ans += strconv.Itoa(res)
		prev[numerator] = len(ans) - 1
		numerator = remain
	}

	return ans
}

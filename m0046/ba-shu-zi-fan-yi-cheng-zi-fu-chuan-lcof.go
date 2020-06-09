package m0046

import "strconv"

// 2020.06.09
// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
// 递归，每次指定跳跃的步数
func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	var translate func(num string)
	count := 0
	translate = func(num string) {
		if len(num) == 0 {
			count++
			return
		}
		if len(num) > 1 && (num[0] == '1' || (num[0] == '2' && num[1] <= '5' && num[1] >= '0')) {
			translate(num[2:])
		}
		translate(num[1:])
	}
	translate(numStr)
	return count
}

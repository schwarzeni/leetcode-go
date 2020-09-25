package _0263

// 2020.09.25
// https://leetcode-cn.com/problems/ugly-number/
func isUgly(num int) bool {
	if num < 6 {
		return num >= 1
	}
	data := []int{2, 3, 5}
	for i := range data {
		for num%data[i] == 0 {
			num /= data[i]
		}
	}
	return num == 1
}

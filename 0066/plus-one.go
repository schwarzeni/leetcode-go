package _0066

// 2020.05.10
// https://leetcode-cn.com/problems/plus-one/
func plusOne(digits []int) []int {
	p, r := 1, make([]int, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		r[i+1], p = (digits[i]+p)%10, (digits[i]+p)/10
	}
	if p == 0 {
		return r[1:]
	}
	r[0] = p
	return r
}

package _12

// 2020.04.11
// https://leetcode-cn.com/problems/integer-to-roman/

var arr = []struct {
	val int
	sig string
}{
	{val: 1000, sig: "M"},
	{val: 900, sig: "CM"},
	{val: 500, sig: "D"},
	{val: 400, sig: "CD"},
	{val: 100, sig: "C"},
	{val: 90, sig: "XC"},
	{val: 50, sig: "L"},
	{val: 40, sig: "XL"},
	{val: 10, sig: "X"},
	{val: 9, sig: "IX"},
	{val: 5, sig: "V"},
	{val: 4, sig: "IV"},
	{val: 1, sig: "I"},
}

func intToRoman(num int) string {
	var result []byte

	for _, v := range arr {
		for num-v.val >= 0 {
			num = num - v.val
			result = append(result, []byte(v.sig)...)
		}
	}

	return string(result)
}

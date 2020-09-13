package mm1626

// 2020.09.13
// https://leetcode-cn.com/problems/calculator-lcci/
func calculate(s string) int {
	var optStack []byte
	var numStack []int

	num := 0
	s += " "
	for idx, b := range s {
		switch {
		case isNumber(b):
			num = num*10 + int(b-'0')
			if idx < len(s)-1 && !isNumber(rune(s[idx+1])) {
				numStack = append(numStack, num)
				num = 0
				for len(optStack) > 0 {
					opt := optStack[len(optStack)-1]
					if opt == '+' || opt == '-' {
						break
					}
					lns := len(numStack)
					v1, v2 := numStack[lns-2], numStack[lns-1]
					res := 0
					if opt == '*' {
						res = v1 * v2
					} else if opt == '/' {
						res = v1 / v2
					}
					optStack = optStack[:len(optStack)-1]
					numStack = numStack[:lns-2]
					numStack = append(numStack, res)
				}
			}
		case isOpt(b):
			optStack = append(optStack, byte(b))
		}
	}
	for len(optStack) > 0 {
		opt := optStack[0]
		v1, v2 := numStack[0], numStack[1]
		res := 0
		if opt == '+' {
			res = v1 + v2
		} else if opt == '-' {
			res = v1 - v2
		}
		optStack = optStack[1:]
		numStack = numStack[1:]
		numStack[0] = res
	}
	return numStack[0]
}

func isNumber(b rune) bool {
	return b >= '0' && b <= '9'
}

func isOpt(b rune) bool {
	return b == '+' || b == '-' || b == '*' || b == '/'
}

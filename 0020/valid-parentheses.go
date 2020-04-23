package _0020

// 2020.04.23
// https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	var (
		stack []int32
		idx   = -1
	)

	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
			idx++
		} else if (idx > -1 && stack != nil) &&
			((v == '}' && stack[idx] == '{') ||
				(v == ']' && stack[idx] == '[') ||
				(v == ')' && stack[idx] == '(')) {
			stack = stack[:len(stack)-1]
			idx--
		} else {
			return false
		}
	}

	return idx == -1
}

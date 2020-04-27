package _0032

// 2020.04.27
// https://leetcode-cn.com/problems/longest-valid-parentheses/
//func longestValidParentheses(s string) int {
//	max := 0
//	idx := 0
//
//	for idx < len(s) {
//		c, i := search(idx, s)
//		if c > max {
//			max = c
//		}
//		if i == idx {
//			idx++
//		} else {
//			idx = i
//		}
//	}
//
//	return max
//}
//
//func search(lidx int, s string) (int, int) {
//	stack := 0
//	count := 0
//	ridx := lidx
//	for idx, ss := range s[lidx:] {
//		if ss == '(' {
//			stack++
//		} else {
//			stack--
//			if stack == 0 {
//				count += idx - ridx + 1 + lidx
//				ridx = idx + 1 + lidx
//			} else if stack < 0 {
//				break
//			}
//		}
//	}
//	return count, ridx
//}

// 使用栈，更快，来自题解，不是我自己想出来的
func longestValidParentheses(s string) int {
	max := 0
	stack := []int{-1}
	for idx, c := range s {
		if c == '(' {
			stack = append(stack, idx)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, idx)
			} else {
				if idx-stack[len(stack)-1] > max {
					max = idx - stack[len(stack)-1]
				}
			}
		}
	}
	return max
}

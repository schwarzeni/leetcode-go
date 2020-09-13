package _0224

// 2020.09.13
// https://leetcode-cn.com/problems/basic-calculator/
var opts = map[rune]int{
	'+': -1,
	'-': -2,
	'*': -3,
	'/': -4,
	'(': -5,
	')': -6,
}

var optFuncs = map[int]func(a, b int) int{
	opts['+']: func(a, b int) int { return a + b }, // -1 --> +
	opts['-']: func(a, b int) int { return a - b }, // -2 --> -
	opts['*']: func(a, b int) int { return a * b }, // -3 --> *
	opts['/']: func(a, b int) int { return a / b }, // -4 --> /
}

var priority = map[int]int{
	opts['+']: 1,
	opts['-']: 1,
	opts['*']: 2,
	opts['/']: 2,
}

func calculate(s string) int {
	var postFixs []int
	stack := Stack{}
	num := 0
	s += " "

	for i, b := range s {
		switch {
		case isNumber(b):
			num = num*10 + int(b-'0')
			if !isNumber(rune(s[i+1])) {
				postFixs = append(postFixs, num)
				num = 0
			}
		case isOpt(b):
			for !stack.Empty() && stack.Peek() != opts['('] &&
				priority[stack.Peek()] <= priority[opts[b]] {
				postFixs = append(postFixs, stack.Pop())
			}
			stack.Push(opts[b])
		case isBracket(b):
			if b == '(' {
				stack.Push(opts[b])
			}
			if b == ')' {
				for !stack.Empty() && stack.Peek() != opts['('] {
					postFixs = append(postFixs, stack.Pop())
				}
				if !stack.Empty() {
					stack.Pop()
				}
			}
		}
	}
	for !stack.Empty() {
		postFixs = append(postFixs, stack.Pop())
	}
	tmpStack := Stack{}
	for _, v := range postFixs {
		if v >= 0 {
			tmpStack.Push(v)
		} else {
			v2 := tmpStack.Pop()
			v1 := tmpStack.Pop()
			tmpStack.Push((optFuncs[v])(v1, v2))
		}
	}
	return tmpStack.Pop()
}

func isNumber(b rune) bool {
	return b >= '0' && b <= '9'
}

func isOpt(b rune) bool {
	return b == '+' || b == '-'
}

func isBracket(b rune) bool {
	return b == '(' || b == ')'
}

type Stack []int

func (s *Stack) Push(d int) {
	*s = append(*s, d)
}

func (s *Stack) Pop() int {
	l := len(*s)
	d := (*s)[l-1]
	*s = (*s)[:l-1]
	return d
}

func (s Stack) Peek() int {
	return s[len(s)-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

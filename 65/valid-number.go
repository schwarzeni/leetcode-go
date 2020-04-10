package _65

// 2020.04.10
// https://leetcode-cn.com/problems/valid-number/
func isNumber(s string) bool {
	const (
		S_INIT     = iota
		S_DECIMAL  // .23
		S_INTEGER  // 23
		S_EXPONENT // e23/ e-1
		S_OPERATOR
	)
	var (
		sLen    = len(s)
		idx     = 0
		isDigit = func(data byte) bool { return data >= '0' && data <= '9' }
		toEnd   = func(idx int) bool {
			for idx < sLen {
				if s[idx] != ' ' {
					return false
				}
				idx += 1
			}
			return true
		}
		status = S_INIT
	)

	for idx < sLen {
		switch status {
		case S_INIT:
			switch {
			case s[idx] == '+' || s[idx] == '-':
				status = S_OPERATOR
				idx += 1
			case isDigit(s[idx]):
				status = S_INTEGER
			case s[idx] == '.':
				status = S_DECIMAL
				idx += 1
			case s[idx] == ' ':
				idx += 1
			default:
				return false
			}
		case S_OPERATOR:
			switch {
			case isDigit(s[idx]):
				status = S_INTEGER
			case s[idx] == '.':
				status = S_DECIMAL
				idx += 1
			default:
				return false
			}
		case S_INTEGER:
			switch {
			case isDigit(s[idx]):
				idx += 1
				if idx == sLen {
					return true
				}
			case s[idx] == 'e':
				status = S_EXPONENT
				idx += 1
			case s[idx] == '.':
				status = S_DECIMAL
				idx += 1
				if idx == sLen {
					return true
				}
			case s[idx] == ' ':
				return toEnd(idx)
			default:
				return false
			}
		case S_EXPONENT:
			switch {
			case isDigit(s[idx]):
				idx += 1
				if idx == sLen {
					return true
				}
			case s[idx] == '+' || s[idx] == '-':
				if s[idx-1] != 'e' {
					return false
				}
				idx += 1
			case s[idx] == ' ':
				if s[idx-1] == 'e' {
					return false
				}
				return toEnd(idx)
			default:
				return false
			}
		case S_DECIMAL:
			switch {
			case isDigit(s[idx]):
				idx += 1
				if idx == sLen {
					return true
				}
			case s[idx] == 'e':
				//  2.2e                    2.e
				if isDigit(s[idx-1]) || (idx-2 >= 0 && isDigit(s[idx-2])) {
					status = S_EXPONENT
					idx += 1
				} else {
					// -.e+43
					return false
				}
			case s[idx] == ' ':
				// fix test_4  ".3 "        "2. "
				if (isDigit(s[idx-1])) || idx-2 >= 0 && isDigit(s[idx-2]) {
					return toEnd(idx)
				}
				return false
			default:
				return false
			}
		}
	}
	return false
}

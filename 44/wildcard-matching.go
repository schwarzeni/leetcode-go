package _44

// 2020.04.03
// https://leetcode-cn.com/problems/wildcard-matching/

func isMatch(s string, p string) bool {
	const (
		T_MATCH_ANY = '*'
		T_MATCH_ONE = '?'
	)
	var (
		match  func(si, pi int) bool
		sLen   = len(s)
		pLen   = len(p)
		record = make([][]int, sLen) // improve cache speed
	)
	for i := 0; i < sLen; i++ {
		record[i] = make([]int, pLen)
	}

	match = func(si, pi int) bool {
		if si == sLen { // fix test_1, test_2
			for _, pp := range p[pi:] {
				if pp != T_MATCH_ANY {
					return false
				}
			}
			return true
		}
		if (si >= sLen && pi < pLen) || (si < sLen && pi >= pLen) {
			return false
		}

		if record[si][pi] == 1 {
			return false
		}
		record[si][pi] = 1

		switch p[pi] {
		case T_MATCH_ONE:
			return match(si+1, pi+1)
		case T_MATCH_ANY:
			// fix test_3
			// fix test_5
			for pi < pLen && p[pi] == T_MATCH_ANY {
				pi += 1
			}
			if pi == pLen {
				return true
			}
			for (si < sLen) && pi < pLen {
				if r := match(si, pi); r {
					return true
				}
				si += 1
			}
			return false
		default:
			if p[pi] != s[si] {
				return false
			}
			return match(si+1, pi+1)
		}
	}
	return match(0, 0)
}

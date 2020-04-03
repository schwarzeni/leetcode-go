package _10

func isMatch(s string, p string) bool {
	const (
		T_MATCH_ANY = '*'
		T_MATCH_ONE = '.'
	)
	var (
		match  func(si, pi int) bool
		sLen   = len(s)
		pLen   = len(p)
		record = make([][]int, sLen+1) // improve cache speed
	)
	for i := 0; i < sLen+1; i++ {
		record[i] = make([]int, pLen+1)
	}

	match = func(si, pi int) bool {
		if si == sLen && pi == pLen {
			return true
		}
		if pi >= pLen {
			return false
		}

		if record[si][pi] == 1 {
			return false
		}
		record[si][pi] = 1

		switch p[pi] {
		case T_MATCH_ONE:
			if pi == pLen-1 && si == sLen-1 {
				return true
			}
			if pi+1 < pLen && p[pi+1] == T_MATCH_ANY {
				return match(si, pi+1)
			}
			if si >= sLen {
				return false
			}
			return match(si+1, pi+1)
		case T_MATCH_ANY:
			if si >= sLen {
				return match(si, pi+1)
			}
			switch p[pi-1] {
			case T_MATCH_ONE:
				if r := match(si, pi+1); r { // fix test_4
					return true
				}
				for si < sLen {
					si += 1 // fix test_5
					if r := match(si, pi+1); r {
						return true
					}
				}
				return false
			default:
				if r := match(si, pi+1); r {
					return true
				}
				for si < sLen && s[si] == p[pi-1] {
					si += 1
					if r := match(si, pi+1); r {
						return true
					}
				}
				return false
			}
		default:
			if pi+1 < pLen && p[pi+1] == T_MATCH_ANY {
				return match(si, pi+1)
			}
			if si >= sLen || p[pi] != s[si] {
				return false
			}
			return match(si+1, pi+1)
		}
	}
	return match(0, 0)
}

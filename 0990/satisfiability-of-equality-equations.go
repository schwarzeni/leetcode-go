package _0990

// 2020.06.08
// https://leetcode-cn.com/problems/satisfiability-of-equality-equations/
func equationsPossible(equations []string) bool {
	var find func(x int) int
	var p = make([]int, 27)
	for i := 1; i < len(p); i++ {
		p[i] = i
	}
	find = func(x int) int {
		if x == p[x] {
			return x
		} else {
			p[x] = find(p[x])
			return p[x]
		}
	}

	for _, eq := range equations {
		if eq[1] == '=' {
			r1 := find(int(eq[0] - 'a'))
			r2 := find(int(eq[3] - 'a'))
			if r1 != r2 {
				p[r2] = r1
			}
		}
	}

	for _, eq := range equations {
		if eq[1] == '!' {
			r1 := find(int(eq[0] - 'a'))
			r2 := find(int(eq[3] - 'a'))
			if r1 == r2 {
				return false
			}
		}
	}

	return true
}
